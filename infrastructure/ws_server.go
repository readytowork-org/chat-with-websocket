package infrastructure

import "boilerplate-api/models"

type WsServer struct {
	Clients    map[*WsClient]bool
	Register   chan *WsClient
	UnRegister chan *WsClient
	Broadcase  chan []byte
	Rooms      map[*models.Room]bool
}

func NewWebscoketServer() *WsServer {
	wsServer := &WsServer{
		Clients:    make(map[*WsClient]bool),
		Register:   make(chan *WsClient),
		UnRegister: make(chan *WsClient),
		Broadcase:  make(chan []byte),
		Rooms:      make(map[*models.Room]bool),
	}
	return wsServer
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
