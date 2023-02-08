package models

type Message struct {
	Base
	Text       string `json:"text"`
	UserRoomId int64  `json:"user_room_id"`
}

type UserMessage struct {
	Message
	User User `json:"user"`
}

func (m Message) TableName() string {
	return "messages"
}
