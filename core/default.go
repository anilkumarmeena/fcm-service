package core

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/fcm-service/config"
	"google.golang.org/api/option"
)

var messagingclient *messaging.Client

func InitFirebase() {
	opt := option.WithCredentialsFile(config.Props.FirebaseJson)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase: %v\n", err)
	}
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	messagingclient = client
}
