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
	uCon := controllers.NewUserController(ser)
	http.HandleFunc("/health", middlewares.CorsMiddleware(controllers.HealthHandler))
	http.HandleFunc("/register", middlewares.CorsMiddleware(aCon.RegisterController))
	http.HandleFunc("/login", middlewares.CorsMiddleware(aCon.LoginController))
	http.HandleFunc("/profile", middlewares.CorsMiddleware(middlewares.AuthMiddleware(uCon.GetProfileController)))
	http.HandleFunc("/update-profile", middlewares.CorsMiddleware(middlewares.AuthMiddleware(uCon.UpdateProfileController)))
	http.ListenAndServe(":8080", nil)
}
