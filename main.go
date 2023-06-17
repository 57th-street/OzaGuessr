package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int
	Message string
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Status: 200, Message: "succeeded"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", nil)
}
