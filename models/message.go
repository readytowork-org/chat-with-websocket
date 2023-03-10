package models

type Message struct {
	Base
	Text   string `json:"text"`
	RoomId int64  `json:"room_id"`
	UserId string `json:"user_id"`
}

type UserMessage struct {
	Message
	User User `json:"user"`
}

func (m Message) TableName() string {
	return "messages"
}
