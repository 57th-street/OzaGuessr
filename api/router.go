package api

import (
	"net/http"

	"github.com/57th-street/oza-gueser/controllers"
	"github.com/57th-street/oza-gueser/api/middlewares"
)

func NewRouter() {
	http.HandleFunc("/health", middlewares.CorsMiddleware(controllers.HealthHandler))
	http.ListenAndServe(":8080", nil)
}
