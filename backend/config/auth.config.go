package config

import (
    "encoding/json"
    "os"
)

type AuthConfig struct {
    Type     string `json:"type"`
    KeyID    string `json:"keyId"`
    Key      string `json:"key"`
    AppID    string `json:"appId"`
    ClientID string `json:"clientId"`
}

func LoadAuthConfig(keyPath string) (*AuthConfig, error) {
    file, err := os.ReadFile(GetEnvConfig(keyPath))
    if err != nil {
        return nil, err
    }

    var config AuthConfig
    if err := json.Unmarshal(file, &config); err != nil {
        return nil, err
    }

    return &config, nil
}