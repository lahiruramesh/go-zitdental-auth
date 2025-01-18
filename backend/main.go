package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"

	"github.com/lahiruramesh/config"
	"github.com/lahiruramesh/middleware"
	"github.com/lahiruramesh/routes"
	
)

func main() {
	config := config.LoadConfig()
	corsHandler := handlers.CORS(
        handlers.AllowedHeaders([]string{
            "Authorization",
            "Content-Type",
            "X-Requested-With",
            "Accept",
            "Origin",
        }),
        handlers.AllowedMethods([]string{
            "GET",
            "POST",
            "PUT",
            "DELETE",
            "OPTIONS",
        }),
        handlers.AllowedOrigins([]string{"http://localhost:5173"}),
        handlers.AllowCredentials(),
        handlers.ExposedHeaders([]string{"Authorization"}),
    )
	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)
	handler := corsHandler(router)
	apiRouter := router.PathPrefix("/api").Subrouter()
	routes.SetupRoutes(apiRouter)

	fmt.Printf("Server is running on port %s in %s mode\n", config.Port, config.Env)
	if err := http.ListenAndServe(":"+config.Port, handler); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
