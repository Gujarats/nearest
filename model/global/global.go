package global

// This file is used for sending a response. used in controller.

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string
	Message string
	Latency float64
	Data    interface{}
}

func SetResponse(w http.ResponseWriter, Status string, Message string) {
	resp := Response{}
	resp.Status = Status
	resp.Message = Message
	json.NewEncoder(w).Encode(resp)
}

func SetResponseTime(w http.ResponseWriter, Status string, Message string, Latency float64) {
	resp := Response{}
	resp.Status = Status
	resp.Message = Message
	resp.Latency = Latency
	json.NewEncoder(w).Encode(resp)
}
