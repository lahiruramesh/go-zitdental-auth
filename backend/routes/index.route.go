package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/lahiruramesh/controller"
	"github.com/lahiruramesh/middleware"
)

func SetupRoutes(r *mux.Router) {
    publicRouters := r.NewRoute().Subrouter()
    publicRouters.HandleFunc("/health", controller.HealthCheckHandler).Methods(http.MethodGet)
    publicRouters.HandleFunc("/allowedUsers", controller.AllowedUsers).Methods(http.MethodGet)
    
    publicRouters.HandleFunc("/oauth/callback", controller.OAuthCallbackHandler).Methods(http.MethodGet)
    publicRouters.HandleFunc("/auth/login", controller.AuthHandler).Methods(http.MethodGet, http.MethodOptions)


    privateRouters := r.NewRoute().Subrouter()
    privateRouters.Use(middleware.CheckAuthentication)
    privateRouters.HandleFunc("/profile", controller.GetUser).Methods(http.MethodGet)
}