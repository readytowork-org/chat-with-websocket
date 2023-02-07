package controllers

import (
	"boilerplate-api/constants"
	"boilerplate-api/models"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type NotifierUser struct {
	models.User
	conn   *websocket.Conn
	server *ChatNotifier
	Send   chan []byte
}

func NewNotifierUser(
	conn *websocket.Conn,
	server *ChatNotifier,
	user models.User,
) *NotifierUser {
	return &NotifierUser{
		User:   user,
		conn:   conn,
		server: server,
		Send:   make(chan []byte),
	}
}
func (client *NotifierUser) readPump() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(constants.MaxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(constants.PongWait))
	client.conn.SetPongHandler(func(string) error {
		client.conn.SetReadDeadline(time.Now().Add(constants.PongWait))
		return nil
	})

	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error : %v", err.Error())
			}
			break
		}

		message := models.Message{}
		err = json.Unmarshal(jsonMessage, &message)
		if err != nil {
			client.server.logger.Zap.Error("Message parsing Error :: ", err.Error())
			return
		}
		message.UserId = client.ID
		client.server.notify <- models.UserMessage{Message: message}
	}
}
func (client *NotifierUser) writePump() {
	ticker := time.NewTicker(constants.PingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.Send:
			client.conn.SetWriteDeadline(time.Now().Add(constants.WriteWait))
			if !ok {
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(client.Send)
			for i := 0; i < n; i++ {
				w.Write(constants.Newline)
				w.Write(<-client.Send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(constants.WriteWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *NotifierUser) disconnect() {
	client.server.unRegister <- client
	close(client.Send)
	client.conn.Close()
}
