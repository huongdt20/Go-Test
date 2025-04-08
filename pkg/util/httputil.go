package pkg

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type baseResponse struct {
	Success     bool              `json:"success"`
	Message     string            `json:"message,omitempty"`
	MessageCode string            `json:"message_code,omitempty"`
	MessageArgs map[string]string `json:"message_args,omitempty"`
	Data        interface{}       `json:"data,omitempty"`
	Detail      interface{}       `json:"detail,omitempty"`
}

func responseJson(w http.ResponseWriter, httpCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	_, _ = w.Write(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
}

func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	responseJson(w, http.StatusOK, &baseResponse{
		Success: true,
		Data:    data,
	})
}

func ResponseError(w http.ResponseWriter, httpCode int, errString string) {
	responseJson(w, httpCode, &baseResponse{
		Success: false,
		Message: errString,
	})
}
