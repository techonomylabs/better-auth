package api

import (
	"better-auth/api/handlers"
	"better-auth/configs"
	"github.com/rs/cors"
	"log"
	"net/http"
)

/**
API ENDPOINTS
CRUD ANY MAP[STRING] INTERFACE IF KEY SET TO TRUE
CRUD USER
*/

func GetMuxAPI() http.Handler {

	log.Print("Initializing Rest Endpoints " + configs.Port)
	mux := http.NewServeMux()

	filteredUserHandler := http.HandlerFunc(handlers.UserHandler)
	//filteredRegisterHandler := http.HandlerFunc(handlers.RegisterHandler)

	mux.Handle("/api/v1/user", handlers.AuthMiddleware(handlers.FilteredMiddleware(filteredUserHandler)))

	mux.HandleFunc("/api/v1/register", handlers.RegisterHandler)
	mux.HandleFunc("/api/v1/user/login", handlers.LoginHandler)
	mux.HandleFunc("/api/v1/user/logout", handlers.LogoutHandler)
	customCors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	handler := customCors.Handler(mux)
	return handler
}
