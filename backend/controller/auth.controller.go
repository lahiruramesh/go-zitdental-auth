package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
    "github.com/lahiruramesh/constants"
	"github.com/lahiruramesh/config"
	"github.com/lahiruramesh/types"
    "github.com/lahiruramesh/utils"
)

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
    basicAuthorization := utils.GetBasicAuthCredentials(clientID, clientSecret)

    data := url.Values{}
    data.Set("grant_type", constants.OAUTH_GRANT_TYPE_AUTHORIZATION_CODE)
    data.Set("code", code)
    data.Set("redirect_uri",authCallBackUrl)
    oauthTokenUrl := utils.GetZitadelURL(constants.OAUTH_TOKEN_PATH)

    headers := map[string]string{
        "Content-Type":  constants.HEADER_URL_ENCODED,
        "Authorization": basicAuthorization,
    }

    req := types.HttpRequest{
        URL:    oauthTokenUrl,
        Method: http.MethodPost,
        Data:   data,
        Headers: headers,
    }

    tokenResp, err := utils.MakeRequest[types.TokenResponse](req)
    if err != nil {
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

func AuthHandler(w http.ResponseWriter, r *http.Request) {
    cf := config.LoadConfig()

    username := r.URL.Query().Get("username")
    authURL := fmt.Sprintf(
        "https://%s/oauth/v2/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&login_hint=%s",
        cf.ZitadelDomain,
        cf.ZitadelCLientID,
        url.QueryEscape(cf.AuthCallbackURL),
        url.QueryEscape("openid profile email"),
        url.QueryEscape(username),
    )
    
    w.Header().Set("Content-Type", constants.HEADER_APPLICATION_JSON)
    json.NewEncoder(w).Encode(map[string]string{
        "redirectUrl": authURL,
    })
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
            Username:    "andrew_s",
            Email:       "andrew_s@gmail.com",
            FirstName:   "Andrew",
            LastName:    "Simon",
            DisplayName: "Andrew Simon",
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

// TODO: Move to a separate file
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
   
    response := types.Response{
        Status:  "success",
        Message: "Service is healthy",
    }
    json.NewEncoder(w).Encode(response)
}