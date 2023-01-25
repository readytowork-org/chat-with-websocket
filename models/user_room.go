package models

type UserRoom struct {
	Base
	UserId string `json:"user_id"`
	RoomId int64  `json:"room_id"`
}

func (r UserRoom) TableName() string {
	return "user_rooms"
}
