package core

import (
	"context"

	"firebase.google.com/go/v4/messaging"
	"github.com/kataras/golog"
)

type TopicMessage struct {
	Data         map[string]string      `json:"data,omitempty"`
	Notification messaging.Notification `json:"notification"`
	Topic        string                 `json:"topic"`
}

func SendMessageTopic(req *TopicMessage) (string, bool) {
	message := &messaging.Message{
		Data:         req.Data,
		Topic:        req.Topic,
		Notification: &req.Notification,
	}
	response, err := messagingclient.Send(context.Background(), message)
	if err != nil {
		golog.Error("Error sending message to topic %s %s \n%s", req.Topic, response, err.Error())
		return "Failed to send message", false
	}
	return "Successfully sent message", true
}
