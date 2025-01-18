package controller

import (
    "encoding/json"
    "net/http"
    "github.com/lahiruramesh/types"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	response := types.Response{
		Status:  "success",
		Message: "User details",
	}
	json.NewEncoder(w).Encode(response)
}