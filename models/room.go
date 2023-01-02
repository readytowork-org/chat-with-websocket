package models

type Room struct {
	Base
	Title string `json:"title"`
}

func (r Room) TableName() string {
	return "rooms"
}
