package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsServer struct {
	chatRoom    map[int64]*Server
	logger      infrastructure.Logger
	roomService services.RoomService
	roomId      chan int64
}

func NewWebSocketServer(logger infrastructure.Logger, roomService services.RoomService) *WsServer {
	return &WsServer{
		logger:      logger,
		chatRoom:    make(map[int64]*Server),
		roomService: roomService,
		roomId:      make(chan int64),
	}
}

type Server struct {
	Clients    map[*WsClient]bool
	Register   chan *WsClient
	UnRegister chan *WsClient
	Broadcase  chan []byte
	Room       models.Room
}

func (w WsServer) ServerWs(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		w.logger.Zap.Error("Error creating upgrader ", err.Error())
		return
	}

	roomId, _ := strconv.ParseInt(c.Param("room-id"), 10, 64)

	room, err := w.roomService.GetRoomWithId(roomId)
	if err != nil {
		w.logger.Zap.Error("No room found", err.Error())
		bytMsg, _ := json.Marshal(gin.H{"msg": "Room Not Found!"})
		conn.WriteMessage(websocket.TextMessage, bytMsg)
		return
	}
	server := w.chatRoom[roomId]
	if server == nil {
		w.chatRoom[roomId] = &Server{
			Clients:    make(map[*WsClient]bool),
			Register:   make(chan *WsClient),
			UnRegister: make(chan *WsClient),
			Broadcase:  make(chan []byte),
			Room:       room,
		}
		w.roomId <- roomId

		server = w.chatRoom[roomId]
	}

	client := NewClient(conn, server)

	go client.writePump()
	go client.readPump()

	server.Register <- client

}

func (w WsServer) RunServer() {

	for {
		select {
		case room := <-w.roomId:
			go w.chatRoom[room].Run()
		}
	}

}

func (server *Server) Run() {
	for {
		select {
		case client := <-server.Register:
			server.RegisterClient(client)
		case client := <-server.UnRegister:
			server.UnRegisterClient(client)
		case message := <-server.Broadcase:
			server.BroadcastToClient(message)
		}
	}

}

func (server *Server) RegisterClient(client *WsClient) {
	server.Clients[client] = true
}

func (server *Server) UnRegisterClient(client *WsClient) {
	if ok := server.Clients[client]; ok {
		delete(server.Clients, client)
	}
}
func (server *Server) BroadcastToClient(message []byte) {
	for client := range server.Clients {
		println("sending msg : ")
		client.Send <- message
	}
}
