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


