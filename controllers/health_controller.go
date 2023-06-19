package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int
	Message string
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: 200, Message: "succeeded"}
	json.NewEncoder(w).Encode(response)
}
