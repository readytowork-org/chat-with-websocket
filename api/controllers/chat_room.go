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
	register       chan *ChatUser
	unRegister     chan *ChatUser
	clients        map[string]*ChatUser
	broadcast      chan models.Message
	messageService services.MessageService
	chatNotifier   *ChatNotifier
}

func NewChatRoom(
	room models.Room,
	logger infrastructure.Logger,
	messageService services.MessageService,
	chatNotifier *ChatNotifier,
) *ChatRoom {
	return &ChatRoom{
		Room:           room,
		clients:        make(map[string]*ChatUser),
		register:       make(chan *ChatUser),
		unRegister:     make(chan *ChatUser),
		broadcast:      make(chan models.Message),
		messageService: messageService,
		logger:         logger,
		chatNotifier:   chatNotifier,
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
			dbMessage, err := chatRoom.messageService.SaveMessageToRoom(models.UserMessage{Message: message})
			if err != nil {
				chatRoom.logger.Zap.Error("Message saving failed", err.Error())
				return
			}
			go func() {
				chatRoom.chatNotifier.notify <- dbMessage
			}()

			bytMsg, _ := json.Marshal(dbMessage)
			chatRoom.BroadcastToClient(bytMsg)
		}
	}
}

func (chatRoom *ChatRoom) RegisterClient(client *ChatUser) {
	chatRoom.clients[client.User.ID] = client
}

func (chatRoom *ChatRoom) UnRegisterClient(client *ChatUser) {
	if ok := chatRoom.clients[client.User.ID]; ok != nil {
		delete(chatRoom.clients, client.User.ID)
	}
}

func (chatRoom *ChatRoom) BroadcastToClient(message []byte) {
	for _, client := range chatRoom.clients {
		client.Send <- message
	}
}
