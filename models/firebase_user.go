package models

type FirebaseAuthUser struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
	Role        string `json:"role"`
	Enabled     int    `json:"enabled"`
	UserId      string `json:"user_id"`
}
