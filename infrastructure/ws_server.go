package infrastructure

import (
	"boilerplate-api/models"

	"github.com/gin-gonic/gin"
)

type WsServer struct {
	Clients    map[*WsClient]bool
	Register   chan *WsClient
	UnRegister chan *WsClient
	Broadcase  chan []byte
	Rooms      map[*models.Room]bool
	logger     Logger
}

func NewWebscoketServer() *WsServer {
	return &WsServer{
		Clients:    make(map[*WsClient]bool),
		Register:   make(chan *WsClient),
		UnRegister: make(chan *WsClient),
		Broadcase:  make(chan []byte),
		Rooms:      make(map[*models.Room]bool),
	}
}

func (w *WsServer) ServerWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		w.logger.Zap.Error("Error creating upgrader ", err.Error())
		return
	}

	client := NewClient(conn, w)

	go client.writePump()
	go client.readPump()

	w.Register <- client
}

func (server *WsServer) Run() {
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

func (server *WsServer) RegisterClient(client *WsClient) {
	server.Clients[client] = true
}

func (server *WsServer) UnRegisterClient(client *WsClient) {
	if ok := server.Clients[client]; ok {
		delete(server.Clients, client)
	}
}
func (server *WsServer) BroadcastToClient(message []byte) {
	for client := range server.Clients {
		client.Send <- message
	}
}
