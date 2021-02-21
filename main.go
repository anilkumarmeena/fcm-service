package main

import (
	"log"
	"net/http"

	"github.com/fcm-service/config"
	"github.com/fcm-service/core"
	"github.com/fcm-service/server"
	"github.com/kataras/golog"
)

func main() {
	if config.Parse("") == false {
		golog.Error("Invalid Config provided")
		return
	}
	core.InitFirebase()
	// core.SendMessageTopic(&core.TopicMessage{Notification: messaging.Notification{Title: "test message"}, Topic: "2"})
	golog.Info("Starting server ", config.Props.ServerAddress)
	handler := http.Handler(http.HandlerFunc(server.RequestHandler))
	if err := http.ListenAndServe(config.Props.ServerAddress, handler); err != nil {
		log.Fatal(err)
	}
	golog.Info("Exiting app...")
}
