package infrastructure

import (
	"context"
	"fmt"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFbApp(logger Logger) (*auth.Client, error) {
	ctx := context.Background()
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKey file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	conf := &firebase.Config{ProjectID: "chat-app-f1246"}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		fmt.Printf("Firebase NewApp (creating firebase app) : %v", err)
	}

	client, errAuth := app.Auth(ctx)
	if errAuth != nil {
		fmt.Print("error initializing firebase auth (creating client)")
	}
	fmt.Print("Firebase app initialized")
	return client, nil
}
