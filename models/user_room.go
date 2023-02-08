package models

type UserRoom struct {
	Id         int64 `json:"id"`
	FollowerId int64 `json:"follower_id"`
	RoomId     int64 `json:"room_id"`
}

func (r UserRoom) TableName() string {
	return "user_rooms"
}
