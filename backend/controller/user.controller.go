package controller

import (
	"encoding/json"
	"net/http"

	"github.com/lahiruramesh/service"
	"github.com/lahiruramesh/types"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(types.TokenContextKey).(string)
	user, err := service.GetProfile(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := types.APISucessResponse[types.Profile]{
		Status:  "success",
		Message: "User profile fetched successfully",
		Results: *user,
	}
	json.NewEncoder(w).Encode(response)
}