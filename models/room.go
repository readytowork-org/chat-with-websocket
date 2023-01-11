package models

type Room struct {
	Base
	Name string `json:"name"`
}

func (r Room) TableName() string {
	return "rooms"
}
