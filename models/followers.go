package models

type Followers struct {
	Base
	UserId       string `json:"user_id"`
	FollowUserId string `json:"follow_user_id"`
}

func (c Followers) TableName() string {
	return "followers"
}
