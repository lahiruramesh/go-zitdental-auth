package routes

import (
    "encoding/json"
	"fmt"
	"net/http"
	"net/url"
    "github.com/gorilla/mux"
	"github.com/lahiruramesh/config"
	"github.com/lahiruramesh/controller"
	"github.com/lahiruramesh/middleware"
)

func SetupRoutes(r *mux.Router) {
    cf := config.LoadConfig()
    publicRouters := r.NewRoute().Subrouter()
    publicRouters.HandleFunc("/health", controller.HealthCheckHandler).Methods("GET")
    publicRouters.HandleFunc("/allowedUsers", controller.AllowedUsers).Methods("GET")
    

    publicRouters.HandleFunc("/oauth/callback", controller.OAuthCallbackHandler).Methods("GET")
    publicRouters.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
        username := r.URL.Query().Get("username")
        authURL := fmt.Sprintf(
            "https://%s/oauth/v2/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&login_hint=%s",
            cf.ZitadelDomain,
            cf.ZitadelCLientID,
            url.QueryEscape(cf.AuthCallbackURL),
            url.QueryEscape("openid profile email"),
            url.QueryEscape(username),
        )
        
        // Return URL instead of redirect
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "redirectUrl": authURL,
        })
    }).Methods("GET", "OPTIONS")


    privateRouters := r.NewRoute().Subrouter()
    privateRouters.Use(middleware.CheckAuthentication)
    
    privateRouters.HandleFunc("/profile", controller.GetUser).Methods("GET")
}