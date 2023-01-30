package models

type Room struct {
	Base
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
	Users     []User `json:"users" gorm:"many2many:user_rooms;"`
}

func (r Room) TableName() string {
	return "rooms"
}
