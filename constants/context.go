package constants

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	// DBTransaction is database transaction handle set at router context
	DBTransaction = "db_trx"

	// Claims -> authentication claims
	Claims = "Claims"

	// UID -> authenticated user's id
	UID = "UID"

	//DUMMYADMIN ->
	DUMMYADMIN = "Administrator"

	//DUMMYEMAIL ->
	DUMMYEMAIL = "dummyrtw@mailinator.com"

	//Adminrole ->
	RoleAdmin = "admin"
)

const (
	// Max wait time when writing message to peer
	WriteWait = 10 * time.Second

	// Max time till next pong from peer
	PongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	PingPeriod = (PongWait * 9) / 10

	// Maximum message size allowed from peer.
	MaxMessageSize = 10000
)

var (
	Newline = []byte{'\n'}
)

var WsUpgrade = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}
