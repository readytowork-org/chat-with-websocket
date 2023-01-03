package models

type Room struct {
	Base
	Name    string `json:"name"`
	OwnerId int    `json:"owner_id"`
}

func (r Room) TableName() string {
	return "rooms"
}
