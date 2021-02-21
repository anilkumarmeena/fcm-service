package server

import (
	"bytes"
	"net/http"
	"runtime/debug"

	"github.com/fcm-service/utils"
	"github.com/kataras/golog"
)

var (
	topic = []byte("/topic")
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("%s\n%s", r, debug.Stack())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
		}
	}()
	if bytes.Equal([]byte(r.Method), []byte("OPTIONS")) {
		w.WriteHeader(http.StatusNoContent) //204
		return
	}
	if utils.PathIs(topic, r) {
		sendtotopic(w, r)
	}
}
