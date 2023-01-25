package models

import "time"

type Followers struct {
	CreatedAt    time.Time `json:"created_at"`
	UserId       string    `json:"user_id"`
	FollowUserId string    `json:"follow_user_id"`
}

func (c Followers) TableName() string {
	return "followers"
}
