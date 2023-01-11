package infrastructure

import (
	"boilerplate-api/models"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

type WsClient struct {
	Client   models.User
	conn     *websocket.Conn
	wsServer *WsServer
	Send     chan []byte
	rooms    map[*models.Room]bool
}

func NewClient(conn *websocket.Conn,
	wsServer *WsServer,
) *WsClient {
	wsClient := &WsClient{
		// Client: models.User{
		// 	Base:     uuid.New(),
		// 	FullName: name,
		// },
		conn:     conn,
		wsServer: wsServer,
		// Send:     make(chan []byte),
		// rooms:    make(map[*models.Room]bool),
	}
	return wsClient
}

func (client *WsClient) readPump() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error {
		client.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, jsonMessage, error := client.conn.ReadMessage()
		if error != nil {
			if websocket.IsUnexpectedCloseError(error, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error : %v", error)
			}
			break
		}
		client.wsServer.Broadcase <- jsonMessage
	}
}
func (client *WsClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.Send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
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
				w.Write(newline)
				w.Write(<-client.Send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
func (client *WsClient) disconnect() {
	client.wsServer.UnRegister <- client
	close(client.Send)
	client.conn.Close()
}
