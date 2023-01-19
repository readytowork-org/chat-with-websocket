package models

type Message struct {
	Base
	Text   string `json:"text"`
	RoomId int64  `json:"room_id"`
	UserId string `json:"user_id"`
}

func (m Message) TableName() string {
	return "messages"
}
