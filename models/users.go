package models

type User struct {
	Base
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type UserWithFollow struct {
	User
	FollowStatus bool `json:"follow_status" `
}

type RoomsUser struct {
	User
	RoomId int64 `json:"-" `
}

// TableName gives table name of model
func (m User) TableName() string {
	return "users"
}
