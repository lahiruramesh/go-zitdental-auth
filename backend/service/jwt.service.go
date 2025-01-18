package service

import (
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "github.com/lahiruramesh/config"
)

type ZitadelClaims struct {
    jwt.RegisteredClaims
}

func CreateZitadelToken() (string, error) {
    // Load auth config
    authConfig, err := config.LoadAuthConfig("ZITADEK_SERVICE_MANAGEMENT_API_KEY_PATH")
    if err != nil {
        return "", fmt.Errorf("failed to load auth config: %w", err)
    }

    // Parse PEM block
    block, _ := pem.Decode([]byte(authConfig.Key))
    if block == nil {
        return "", fmt.Errorf("failed to parse PEM block containing private key")
    }

    // Parse RSA private key
    privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return "", fmt.Errorf("failed to parse private key: %w", err)
    }
    // Create claims
    claims := ZitadelClaims{
        RegisteredClaims: jwt.RegisteredClaims{
            Issuer:    authConfig.ClientID,
            Subject:   authConfig.ClientID,
            Audience:  jwt.ClaimStrings{"https://myinstance-3tqpfz.us1.zitadel.cloud"},
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    // Create token with headers
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    token.Header["kid"] = authConfig.KeyID
	token.Header["alg"] = jwt.SigningMethodRS256.Name

    // Sign token
    signedToken, err := token.SignedString(privateKey)
    if err != nil {
        return "", fmt.Errorf("failed to sign token: %w", err)
    }

    return signedToken, nil
}