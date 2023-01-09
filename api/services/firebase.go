package services

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"context"

	"firebase.google.com/go/auth"
)

type FirebaseService struct {
	fbAuth *auth.Client
	logger infrastructure.Logger
}

func NewFirebaseService(
	fbAuth *auth.Client,
	logger infrastructure.Logger,
) FirebaseService {
	return FirebaseService{
		fbAuth: fbAuth,
		logger: logger,
	}
}

func (c FirebaseService) CreateUser(newUser models.User) (*auth.UserRecord, error) {
	user := (&auth.UserToCreate{}).Email(newUser.Email).DisplayName(newUser.FullName)
	record, err := c.fbAuth.CreateUser(context.Background(), user)
	if err != nil {
		c.logger.Zap.Info(err)
	}

	return record, err
}
