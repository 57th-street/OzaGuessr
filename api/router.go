package api

import (
	"database/sql"
	"net/http"

	"github.com/57th-street/oza-gueser/api/middlewares"
	"github.com/57th-street/oza-gueser/controllers"
	"github.com/57th-street/oza-gueser/services"
)

func NewRouter(db *sql.DB) {
	ser := services.NewService(db)
	aCon := controllers.NewAuthController(ser)
	http.HandleFunc("/health", middlewares.CorsMiddleware(controllers.HealthHandler))
	http.HandleFunc("/register", middlewares.CorsMiddleware(aCon.RegisterController))
	http.HandleFunc("/login", middlewares.CorsMiddleware(aCon.LoginController))
	http.ListenAndServe(":8080", nil)
}
