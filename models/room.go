package models

type Room struct {
	Base
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
}

type RoomWithUsers struct {
	Room
	Users []User `json:"users" gorm:"many2many:user_rooms;joinForeignKey:room_id"`
}

func (r Room) TableName() string {
	return "rooms"
}
