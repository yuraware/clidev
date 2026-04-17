package auth

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// AppleJWT generates a bearer token for the Apple App Store Connect API.
// It uses an ES256 private key (.p8 file) plus key ID and issuer ID.
type AppleJWT struct {
	KeyID      string
	IssuerID   string
	PrivateKey *ecdsa.PrivateKey
	Audience   string
	Expiry     time.Duration
}

// NewAppleJWT loads credentials from environment variables or file paths.
//
// Required config keys: key_id, issuer_id, private_key (env or file).
// Optional: audience (default "appstoreconnect-v1"), expiry_seconds (default 1200).
func NewAppleJWT(cfg map[string]resolvedValue) (*AppleJWT, error) {
	keyID := cfg["key_id"].String()
	issuerID := cfg["issuer_id"].String()
	keyData := cfg["private_key"].String()

	if keyID == "" {
		return nil, fmt.Errorf("apple_jwt: missing key_id")
	}
	if issuerID == "" {
		return nil, fmt.Errorf("apple_jwt: missing issuer_id")
	}
	if keyData == "" {
		return nil, fmt.Errorf("apple_jwt: missing private_key")
	}

	pk, err := parseECPrivateKey([]byte(keyData))
	if err != nil {
		return nil, fmt.Errorf("apple_jwt: parsing private key: %w", err)
	}

	audience := "appstoreconnect-v1"
	if a := cfg["audience"].String(); a != "" {
		audience = a
	}

	expiry := 20 * time.Minute
	// expiry_seconds could be added here if needed.

	return &AppleJWT{
		KeyID:      keyID,
		IssuerID:   issuerID,
		PrivateKey: pk,
		Audience:   audience,
		Expiry:     expiry,
	}, nil
}

// AuthHeader implements Provider by generating a fresh JWT bearer token.
func (a *AppleJWT) AuthHeader() (string, error) {
	token, err := a.Token()
	if err != nil {
		return "", err
	}
	return "Bearer " + token, nil
}

// Token generates a fresh JWT.
func (a *AppleJWT) Token() (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"iss": a.IssuerID,
		"iat": now.Unix(),
		"exp": now.Add(a.Expiry).Unix(),
		"aud": a.Audience,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = a.KeyID

	signed, err := token.SignedString(a.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("signing jwt: %w", err)
	}
	return signed, nil
}

func parseECPrivateKey(data []byte) (*ecdsa.PrivateKey, error) {
	// Support raw PEM or base64-like content in an env var.
	// If the content doesn't look like PEM, try loading it as a file path.
	if len(data) > 0 && data[0] != '-' {
		fileData, err := os.ReadFile(string(data))
		if err == nil {
			data = fileData
		}
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("no PEM block found")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		// Fallback: try ParseECPrivateKey
		ecKey, err2 := x509.ParseECPrivateKey(block.Bytes)
		if err2 != nil {
			return nil, fmt.Errorf("ParsePKCS8: %v; ParseEC: %v", err, err2)
		}
		return ecKey, nil
	}

	ecKey, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("key is not ECDSA (got %T)", key)
	}
	return ecKey, nil
}
