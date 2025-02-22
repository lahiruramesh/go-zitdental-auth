package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
    "context"

	"github.com/lahiruramesh/service"
    "github.com/lahiruramesh/constants"
    "github.com/lahiruramesh/types"
)

func CheckAuthentication(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            w.Header().Set("Content-Type", constants.HEADER_APPLICATION_JSON)
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(map[string]string{
                "error": "Unauthorized: No token provided",
            })
            return
        }

        tokenString := strings.Replace(token, "Bearer ", "", 1)

        isValid, err := service.VerifyToken(tokenString)
        if err != nil || !isValid {
            w.Header().Set("Content-Type", constants.HEADER_APPLICATION_JSON)
            w.WriteHeader(http.StatusForbidden)
            json.NewEncoder(w).Encode(map[string]string{
                "error": "Forbidden: Invalid token",
            })
            return
        }

        ctx := context.WithValue(r.Context(), types.TokenContextKey, tokenString)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}