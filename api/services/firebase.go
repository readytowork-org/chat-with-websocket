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

func (fb *FirebaseService) VerifyToken(idToken string) (*auth.Token, error) {
	token, err := fb.fbAuth.VerifyIDToken(context.Background(), idToken)
	return token, err
}
func (c FirebaseService) CreateUser(newUser models.FirebaseAuthUser) (*auth.UserRecord, error) {
	user := (&auth.UserToCreate{}).Email(newUser.Email).DisplayName(newUser.DisplayName)
	record, err := c.fbAuth.CreateUser(context.Background(), user)
	if err != nil {
		c.logger.Zap.Info(err)
	}
	claims := map[string]interface{}{
		"role":   newUser.Role,
		"fb_uid": record.UID,
		"id":     newUser.UserId,
	}
	err = c.fbAuth.SetCustomUserClaims(context.Background(), record.UID, claims)
	if err != nil {
		return nil, err
	}
	return record, err
}
