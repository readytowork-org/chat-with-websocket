package models

type UserRoom struct {
	Base
	UserId string `json:"user_id"`
	RoomId int64  `json:"room_id"`
	IsPrivate bool `json:"is_private"`
}

func (r UserRoom) TableName() string {
	return "user_rooms"
}
