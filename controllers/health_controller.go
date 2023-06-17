package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status  int
	Message string
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("access")
	response := Response{Status: 200, Message: "succeeded"}
	json.NewEncoder(w).Encode(response)
}
