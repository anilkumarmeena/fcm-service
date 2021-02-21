package server

import (
	"net/http"

	"github.com/fcm-service/core"
	"github.com/fcm-service/utils"
)

func sendtotopic(w http.ResponseWriter, r *http.Request) {
	if r.Method != utils.HTTPPOST {
		utils.SendJSONResponse(w, "Method not allowed", http.StatusNotFound)
		return
	}
	var reqBody core.TopicMessage
	err := utils.GetRequestBody(r, &reqBody)
	if err != nil {
		utils.SendJSONResponse(w, "Invalid Notification Body", http.StatusBadRequest)
		return
	}
	go core.SendMessageTopic(&reqBody)
	// msg, sent := core.SendMessageTopic(&reqBody)
	// if !sent {
	// 	utils.SendJSONResponse(w, msg, http.StatusBadRequest)
	// 	return
	// }
	utils.SendJSONResponse(w, "Message Sent", http.StatusOK)
}
