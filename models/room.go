package models

type Room struct {
	Base
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
}

func (r Room) TableName() string {
	return "rooms"
}
