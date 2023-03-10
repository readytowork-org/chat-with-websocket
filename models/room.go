package models

type Room struct {
	Base
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
}

type RoomWithUsers struct {
	Room
	Users []RoomsUser `json:"users" gorm:"foreignKey:room_id"`
}

func (r Room) TableName() string {
	return "rooms"
}
