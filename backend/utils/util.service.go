package utils

import (
    "encoding/json"
    "encoding/base64"
    "fmt"
    "net/http"
    "net/url"
    "strings"

	"github.com/lahiruramesh/types"
    "github.com/lahiruramesh/config"
)

func MakeRequest[T any](req types.HttpRequest) (*T, error) {
    httpReq, err := http.NewRequest(
        req.Method,
        req.URL,
        strings.NewReader(req.Data.Encode()),
    )
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

	for key, value := range req.Headers {
        httpReq.Header.Add(key, value)
    }

    client := &http.Client{}
    resp, err := client.Do(httpReq)
    if err != nil {
        return nil, fmt.Errorf("request failed: %w", err)
    }
    defer resp.Body.Close()

    var tokenResp T
    if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    return &tokenResp, nil
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