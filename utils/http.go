package utils

import (
	"encoding/json"
	"net/http"
)

const (
	HTTPGET    = "GET"
	HTTPPOST   = "POST"
	HTTPPUT    = "OUT"
	HTTPDELETE = "DELETE"
)

func GetRequestBody(r *http.Request, sourceRef interface{}) error {
	return json.NewDecoder(r.Body).Decode(&sourceRef)
}

func GetQueryParams(r *http.Request) map[string][]string {
	return r.URL.Query()
}

func SendJSONResponse(w http.ResponseWriter, msg string, code int, extras ...int) {
	var httpStatus int
	if len(extras) > 0 {
		httpStatus = extras[0]
	} else {
		httpStatus = code
	}
	w.WriteHeader(httpStatus)
	w.Header().Add("Content-Type", "application/json")
	res := responseBody{msg, code}
	resByte, _ := json.Marshal(&res)
	w.Write(resByte)
}
