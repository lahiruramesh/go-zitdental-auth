package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/lahiruramesh/config"
	"github.com/lahiruramesh/service"
	"github.com/lahiruramesh/types"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
   
    response := types.Response{
        Status:  "success",
        Message: "Service is healthy",
    }
    json.NewEncoder(w).Encode(response)
}

type TokenResponse struct {
    AccessToken  string `json:"access_token"`
    TokenType    string `json:"token_type"`
    ExpiresIn    int    `json:"expires_in"`
    RefreshToken string `json:"refresh_token"`
    IdToken      string `json:"id_token"`
}


func OAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
    cf := config.LoadConfig()
    code := r.URL.Query().Get("code")

    if code == "" {
        http.Error(w, "Missing authorization code", http.StatusBadRequest)
        return
    }
    clientID := cf.ZitadelCLientID
    clientSecret := cf.ZitadelClientSecret
    authCallBackUrl := cf.AuthCallbackURL
    basicAuthorization := service.GetBasicAuthCredentials(clientID, clientSecret)

    data := url.Values{}
    data.Set("grant_type", "authorization_code")
    data.Set("code", code)
    data.Set("redirect_uri",authCallBackUrl)
    oauthTokenUrl := service.GetZitadelURL("oauth/v2/token")

    req, err := http.NewRequest("POST", 
        oauthTokenUrl,
        strings.NewReader(data.Encode()))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Authorization", basicAuthorization)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var tokenResp TokenResponse
    if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    frontendCallback := fmt.Sprintf(
        "%s?token=%s",
        cf.WebCallbackURL,
        tokenResp.AccessToken,
    )
    http.Redirect(w, r, frontendCallback, http.StatusTemporaryRedirect)
}


func AllowedUsers(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement the function to get users from roles from Auth Service
    // Original Implementation planned to get users from roles from Auth Service 
    // by using the ZitadelGetUserByRoles() function from the zitadel service
    // and return the response to the client
    // The implementation is not completed yet
    // Security Check passing sensitive data with masking
    users := []types.UserProfile{
        {
            ID:          "user1",
            Username:    "ab",
            Email:       "john.doe@example.com",
            FirstName:   "John",
            LastName:    "Doe",
            DisplayName: "John Doe",
            Roles:       []string{"user", "admin"},
        },
        {
            ID:          "user2",
            Username:    "ab",
            Email:       "jane.smith@example.com",
            FirstName:   "Jane",
            LastName:    "Smith",
            DisplayName: "Jane Smith",
            Roles:       []string{"user"},
        },
        {
            ID:          "user3",
            Username:    "ab",
            Email:       "jane.smith@example.com",
            FirstName:   "James",
            LastName:    "Bond",
            DisplayName: "James Bond",
            Roles:       []string{"user"},
        },
    }

    w.Header().Set("Content-Type", "application/json")
    response := types.APISucessResponse[[]types.UserProfile]{
        Status:  "success",
        Message: "Users retrieved successfully",
        Results:    users,
    }
    
    json.NewEncoder(w).Encode(response)   
}