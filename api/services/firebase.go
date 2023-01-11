package services

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"context"

	"firebase.google.com/go/auth"
)

type FirebaseService struct {
	auth   *auth.Client
	logger infrastructure.Logger
}

func NewFirebaseService(
	fbAuth *auth.Client,
	logger infrastructure.Logger,
) FirebaseService {
	return FirebaseService{
		auth:   fbAuth,
		logger: logger,
	}
}

func (fb *FirebaseService) VerifyToken(idToken string) (*auth.Token, error) {
	token, err := fb.auth.VerifyIDToken(context.Background(), idToken)
	return token, err
}

func (fb *FirebaseService) CreateUser(newUser models.FirebaseAuthUser) (*auth.UserRecord, error) {
	user := (&auth.UserToCreate{}).Email(newUser.Email).DisplayName(newUser.FullName).Password(newUser.Password)
	record, err := fb.auth.CreateUser(context.Background(), user)
	if err != nil {
		fb.logger.Zap.Info(err)
	}
	return record, err
}
