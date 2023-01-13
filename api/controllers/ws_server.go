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
	server      map[int]*Server
	logger      infrastructure.Logger
	roomService services.RoomService
}

func NewWebSocketServer(logger infrastructure.Logger, roomService services.RoomService) *WsServer {
	return &WsServer{
		logger:      logger,
		server:      make(map[int]*Server),
		roomService: roomService,
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
	server := w.server[int(roomId)]
	if server == nil {
		server = &Server{
			Clients:    make(map[*WsClient]bool),
			Register:   make(chan *WsClient),
			UnRegister: make(chan *WsClient),
			Broadcase:  make(chan []byte),
			Room:       room,
		}
	}

	client := NewClient(conn, server)

	go client.writePump()
	go client.readPump()

	server.Register <- client
}

func (w WsServer) RunServer() {
	for _, server := range w.server {
		server.Run()
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
		client.Send <- message
	}
}
