package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/constants"
	"boilerplate-api/infrastructure"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
)

type WsServer struct {
	servers        map[int64]*ChatRoom
	newRoomId      chan int64
	logger         infrastructure.Logger
	roomService    services.RoomService
	userService    services.UserService
	messageService services.MessageService
}

func NewWebSocketServer(
	logger infrastructure.Logger,
	roomService services.RoomService,
	userService services.UserService,
	messageService services.MessageService,
) *WsServer {
	return &WsServer{
		logger:         logger,
		servers:        make(map[int64]*ChatRoom),
		newRoomId:      make(chan int64),
		roomService:    roomService,
		userService:    userService,
		messageService: messageService,
	}
}

func (w *WsServer) ServerWs(c *gin.Context) {
	conn, err := wsUpgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		w.logger.Zap.Error("Error creating wsUpgrade ", err.Error())
		return
	}

	roomId, _ := strconv.ParseInt(c.Param("room-id"), 10, 64)
	room, err := w.roomService.GetRoomById(roomId)
	if err != nil {
		w.logger.Zap.Error("No room found", err.Error())
		bytMsg, _ := json.Marshal(gin.H{"msg": "room Not Found!"})
		err = conn.WriteMessage(websocket.TextMessage, bytMsg)
		if err != nil {
			w.logger.Zap.Error("WriteMessage", err.Error())
			return
		}
		return
	}

	userId := c.MustGet(constants.UID).(string)
	user, err := w.userService.GetOneUserById(userId)
	if err != nil {
		w.logger.Zap.Error("No user found", err.Error())
		bytMsg, _ := json.Marshal(gin.H{"msg": "User Not Found!"})
		err = conn.WriteMessage(websocket.TextMessage, bytMsg)
		if err != nil {
			w.logger.Zap.Error("WriteMessage", err.Error())
			return
		}
		return
	}

	server := w.servers[roomId]
	if server == nil {
		w.servers[roomId] = NewChatRoom(room, w.logger, w.messageService)
		server = w.servers[roomId]
		w.newRoomId <- roomId
	}

	client := NewClient(conn, server, user)

	go client.writePump()
	go client.readPump()

	server.register <- client
}

func (w *WsServer) RunServer() {
	for {
		select {
		case roomId := <-w.newRoomId:
			go w.servers[roomId].Run()
		}
	}
}
