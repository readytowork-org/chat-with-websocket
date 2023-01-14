package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"encoding/json"
)

type ChatRoom struct {
	models.Room
	logger         infrastructure.Logger
	register       chan *WsClient
	unRegister     chan *WsClient
	clients        map[string]*WsClient
	broadcast      chan models.Message
	messageService services.MessageService
}

func NewChatRoom(
	room models.Room,
	logger infrastructure.Logger,
	messageService services.MessageService,
) *ChatRoom {
	return &ChatRoom{
		Room:           room,
		clients:        make(map[string]*WsClient),
		register:       make(chan *WsClient),
		unRegister:     make(chan *WsClient),
		broadcast:      make(chan models.Message),
		messageService: messageService,
		logger:         logger,
	}
}

func (chatRoom *ChatRoom) Run() {
	for {
		select {
		case client := <-chatRoom.register:
			chatRoom.RegisterClient(client)
		case client := <-chatRoom.unRegister:
			chatRoom.UnRegisterClient(client)
		case message := <-chatRoom.broadcast:
			dbMessage, err := chatRoom.messageService.SaveMessageToRoom(message)
			if err != nil {
				chatRoom.logger.Zap.Error("No user found", err.Error())
				return
			}
			bytMsg, _ := json.Marshal(dbMessage)
			chatRoom.BroadcastToClient(bytMsg)
		}
	}
}

func (chatRoom *ChatRoom) RegisterClient(client *WsClient) {
	chatRoom.clients[client.User.ID] = client
}

func (chatRoom *ChatRoom) UnRegisterClient(client *WsClient) {
	if ok := chatRoom.clients[client.User.ID]; ok != nil {
		delete(chatRoom.clients, client.User.ID)
	}
}

func (chatRoom *ChatRoom) BroadcastToClient(message []byte) {
	for _, client := range chatRoom.clients {
		client.Send <- message
	}
}
