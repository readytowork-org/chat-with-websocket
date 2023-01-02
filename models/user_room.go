package models

type UserRoom struct {
	Base
	UserId int64 `json:"user_id"`
	RoomId int64 `json:"room_id"`
}

func (r UserRoom) TableName() string {
	return "user_rooms"
}
