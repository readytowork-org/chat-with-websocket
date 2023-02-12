package models

type Room struct {
	Base
	Name       string `json:"name"`
	FollowerId int64  `json:"follower_id"`
	IsPrivate  bool   `json:"is_private"`
}

type RoomWithUsers struct {
	Room
	Users []RoomsUser `json:"users" gorm:"foreignKey:room_id"`
}

func (r Room) TableName() string {
	return "rooms"
}
