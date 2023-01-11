package models

type FirebaseAuthUser struct {
	User
	Password string `json:"password"`
}
