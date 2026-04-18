package auth

import (
	"encoding/base64"
	"fmt"
)

type basicProvider struct {
	username string
	password string
}

func (b *basicProvider) AuthHeader() (string, error) {
	if b.username == "" {
		return "", fmt.Errorf("basic auth: missing username")
	}
	creds := base64.StdEncoding.EncodeToString([]byte(b.username + ":" + b.password))
	return "Basic " + creds, nil
}
