package service

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/lahiruramesh/config"
	"github.com/lahiruramesh/constants"
	"github.com/lahiruramesh/types"
	"github.com/lahiruramesh/utils"
)


func VerifyToken(tokenString string) (bool, error) {
    cf := config.LoadConfig()
	clientID :=cf.ZitadelCLientID
    clientSecret := cf.ZitadelClientSecret
	oauthIntrospectionUrl := utils.GetZitadelURL(constants.OAUTH_INTROSPECT_PATH)
    basicAuthorization := utils.GetBasicAuthCredentials(clientID, clientSecret)
    
    data := url.Values{}
    data.Set("token", strings.TrimPrefix(tokenString, "Bearer "))

    headers := map[string]string{
        "Content-Type":  constants.HEADER_URL_ENCODED,
        "Authorization": basicAuthorization,
    }    
    req := types.HttpRequest{
        URL:    oauthIntrospectionUrl,
        Method: http.MethodPost,
        Data:   data,
        Headers: headers,
    }

    resp, err := utils.MakeRequest[types.IntrospectionResponse](req)
    if err != nil {
        return false, fmt.Errorf("failed to make request: %w", err)
    }
    
    return resp.Active, nil
}

func GetProfile(accessToken string) (*types.Profile, error) {
    headers := map[string]string{
        "Content-Type":  constants.HEADER_APPLICATION_JSON,
        "Accept": constants.HEADER_APPLICATION_JSON,
        "Authorization": "Bearer " + accessToken,
    }
    url := utils.GetZitadelURL(constants.OAUTH_USERINFO_PATH)

    req := types.HttpRequest{
        URL:     url,
        Method:  http.MethodGet,
        Headers: headers,
    }

    profile, err := utils.MakeRequest[types.Profile](req)
    if err != nil {
        return nil, fmt.Errorf("failed to get user profile: %w", err)
    }

    return profile, nil
}