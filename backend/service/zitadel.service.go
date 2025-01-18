package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/lahiruramesh/config"
	"github.com/lahiruramesh/types"
)


func VerifyToken(tokenString string) (bool, error) {
    cf := config.LoadConfig()
	clientID :=cf.ZitadelCLientID
    clientSecret := cf.ZitadelClientSecret
	oauthIntrospectionUrl := GetZitadelURL("oauth/v2/introspect")
    
    basicAuthorization := GetBasicAuthCredentials(clientID, clientSecret)
    
    data := url.Values{}
    data.Set("token", strings.TrimPrefix(tokenString, "Bearer "))
    
    req, err := http.NewRequest(
        "POST",
        oauthIntrospectionUrl,
        strings.NewReader(data.Encode()),
    )
    if err != nil {
        return false, fmt.Errorf("failed to create request: %w", err)
    }
    
    // Set headers
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Authorization", "Basic "+basicAuthorization)
    
    // Make request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return false, fmt.Errorf("introspection request failed: %w", err)
    }
    defer resp.Body.Close()
    
    // Parse response
    var introspection types.IntrospectionResponse
    if err := json.NewDecoder(resp.Body).Decode(&introspection); err != nil {
        return false, fmt.Errorf("failed to decode response: %w", err)
    }
    
    return introspection.Active, nil
}


func GetZitadelURL(path string) string {
	return fmt.Sprintf("https://%s/%s", config.LoadConfig().ZitadelDomain, path)
}

func GetBasicAuthCredentials(clientID string, clientSecret string) string {
    credentials := fmt.Sprintf("%s:%s", 
    url.QueryEscape(clientID), 
    url.QueryEscape(clientSecret))
    basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials))
    return basicAuth
}