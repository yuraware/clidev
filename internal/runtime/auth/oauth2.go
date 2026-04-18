package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// oauth2Provider implements the client-credentials OAuth 2.0 flow.
// It caches the token and refreshes automatically when it expires.
type oauth2Provider struct {
	clientID     string
	clientSecret string
	tokenURL     string
	scopes       string

	mu        sync.Mutex
	token     string
	expiresAt time.Time
}

func (o *oauth2Provider) AuthHeader() (string, error) {
	token, err := o.getToken()
	if err != nil {
		return "", err
	}
	return "Bearer " + token, nil
}

func (o *oauth2Provider) getToken() (string, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.token != "" && time.Now().Before(o.expiresAt) {
		return o.token, nil
	}

	form := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {o.clientID},
		"client_secret": {o.clientSecret},
	}
	if o.scopes != "" {
		form.Set("scope", o.scopes)
	}

	resp, err := http.PostForm(o.tokenURL, form)
	if err != nil {
		return "", fmt.Errorf("oauth2 token request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading token response: %w", err)
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("oauth2 token error %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("parsing token response: %w", err)
	}

	o.token = result.AccessToken
	if result.ExpiresIn > 0 {
		o.expiresAt = time.Now().Add(time.Duration(result.ExpiresIn-30) * time.Second)
	} else {
		o.expiresAt = time.Now().Add(55 * time.Minute)
	}

	return o.token, nil
}

// newOAuth2Provider creates a client-credentials provider.
func newOAuth2Provider(resolved map[string]resolvedValue) (Provider, error) {
	clientID := resolved["client_id"].String()
	clientSecret := resolved["client_secret"].String()
	tokenURL := resolved["token_url"].String()

	if clientID == "" {
		// Support simple bearer token fallback (e.g. gcloud access token).
		if t := resolved["token"].String(); t != "" {
			return &bearerProvider{token: t}, nil
		}
		return nil, fmt.Errorf("oauth2: missing client_id (or token for direct bearer)")
	}
	if tokenURL == "" {
		return nil, fmt.Errorf("oauth2: missing token_url")
	}

	scopes := resolved["scopes"].String()
	// Normalise scope separators.
	scopes = strings.ReplaceAll(scopes, ",", " ")

	return &oauth2Provider{
		clientID:     clientID,
		clientSecret: clientSecret,
		tokenURL:     tokenURL,
		scopes:       scopes,
	}, nil
}
