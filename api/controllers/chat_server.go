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

type ChatServer struct {
	servers        map[int64]*ChatRoom
	newRoomId      chan int64
	deleteRoomId   chan int64
	chatNotifier   *ChatNotifier
	logger         infrastructure.Logger
	roomService    services.RoomService
	userService    services.UserService
	messageService services.MessageService
}

func NewChatServer(
	logger infrastructure.Logger,
	roomService services.RoomService,
	userService services.UserService,
	messageService services.MessageService,
	chatNotifier *ChatNotifier,
) *ChatServer {
	return &ChatServer{
		logger:         logger,
		chatNotifier:   chatNotifier,
		servers:        make(map[int64]*ChatRoom),
		newRoomId:      make(chan int64),
		deleteRoomId:   make(chan int64),
		roomService:    roomService,
		userService:    userService,
		messageService: messageService,
	}
}

func (w *ChatServer) ServerWs(c *gin.Context) {
	roomId, _ := strconv.ParseInt(c.Param("room-id"), 10, 64)
	userId := c.MustGet(constants.UID).(string)

	conn, err := constants.WsUpgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		w.logger.Zap.Error("Error creating wsUpgrade ", err.Error())
		return
	}

	server := w.servers[roomId]
	if server == nil {
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
		w.servers[roomId] = NewChatRoom(room, w.logger, w.messageService, w.chatNotifier)
		server = w.servers[roomId]
		w.newRoomId <- roomId
	}

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

	client := NewChatUser(conn, server, user)

	go client.writePump()
	go client.readPump()

	server.register <- client
}

func (w *ChatServer) RunServer() {
	for {
		select {
		case roomId := <-w.newRoomId:
			go w.servers[roomId].Run()
		case roomId := <-w.deleteRoomId:
			delete(w.servers, roomId)
		}
	}
}
