package global

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string
	Message string
	Data    interface{}
}

func SetResponse(w http.ResponseWriter, Status string, Message string) {
	resp := Response{}
	resp.Status = Status
	resp.Message = Message
	json.NewEncoder(w).Encode(resp)
}
