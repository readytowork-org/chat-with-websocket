package models

type Message struct {
	Base
	Text   string `json:"text"`
	RoomId int64  `json:"room_id"`
	UserId int64  `json:"user_id"`
}

func (m Message) TableName() string {
	return "messages"
}
