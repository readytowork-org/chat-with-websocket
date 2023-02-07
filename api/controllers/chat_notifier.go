package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/constants"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatNotifier struct {
	logger      infrastructure.Logger
	userService services.UserService
	clients     map[string]*NotifierUser
	notify      chan models.UserMessage
	register    chan *NotifierUser
	unRegister  chan *NotifierUser
}

func NewChatNotifier(
	logger infrastructure.Logger,
	userService services.UserService,
) *ChatNotifier {
	return &ChatNotifier{
		logger:      logger,
		userService: userService,
		clients:     make(map[string]*NotifierUser),
		notify:      make(chan models.UserMessage),
		register:    make(chan *NotifierUser),
		unRegister:  make(chan *NotifierUser),
	}
}

func (chatNotifier *ChatNotifier) ServerWs(c *gin.Context) {
	userId := c.MustGet(constants.UID).(string)

	conn, err := constants.WsUpgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		chatNotifier.logger.Zap.Error("Error creating wsUpgrade ", err.Error())
		return
	}

	server := chatNotifier.clients[userId]
	if server == nil {
		user, err := chatNotifier.userService.GetOneUserById(userId)
		if err != nil {
			chatNotifier.logger.Zap.Error("No user found", err.Error())
			bytMsg, _ := json.Marshal(gin.H{"msg": "User Not Found!"})
			err = conn.WriteMessage(websocket.TextMessage, bytMsg)
			if err != nil {
				chatNotifier.logger.Zap.Error("WriteMessage", err.Error())
				return
			}
			return
		}
		client := NewNotifierUser(conn, chatNotifier, user)

		go client.writePump()
		go client.readPump()

		chatNotifier.register <- client
	}
}

func (chatNotifier *ChatNotifier) RunServer() {
	for {
		select {
		case client := <-chatNotifier.register:
			chatNotifier.RegisterClient(client)
		case client := <-chatNotifier.unRegister:
			chatNotifier.UnRegisterClient(client)
		case message := <-chatNotifier.notify:
			chatNotifier.BroadcastToClient(message)
		}
	}
}

func (chatNotifier *ChatNotifier) RegisterClient(client *NotifierUser) {
	chatNotifier.clients[client.User.ID] = client
}

func (chatNotifier *ChatNotifier) UnRegisterClient(client *NotifierUser) {
	if ok := chatNotifier.clients[client.User.ID]; ok != nil {
		delete(chatNotifier.clients, client.User.ID)
	}
}

func (chatNotifier *ChatNotifier) BroadcastToClient(message models.UserMessage) {
	users, err := chatNotifier.userService.GetUsersByRoomId(message.RoomId, message.UserId)
	if err != nil {
		chatNotifier.logger.Zap.Error("Failed to get room users", err.Error())
		return
	}
	for _, user := range users {
		client := chatNotifier.clients[user.ID]
		if client != nil {
			bytMsg, _ := json.Marshal(message)
			client.Send <- bytMsg
		}
	}
}
