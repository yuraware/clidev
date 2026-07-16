package auth

import (
	"fmt"
	"os"

	"github.com/yuraware/clidev/internal/cliSchema"
)

// Provider can attach auth to an HTTP request header.
type Provider interface {
	// AuthHeader returns the "Authorization" header value (e.g. "Bearer <token>").
	AuthHeader() (string, error)
}

// resolvedValue is a string that was resolved from an EnvOrValue config entry.
type resolvedValue struct {
	v string
}

func (r resolvedValue) String() string { return r.v }

// New creates the appropriate Provider for the given AuthConfig.
//
// Supported types:
//
//	none        — no auth
//	bearer      — Authorization: Bearer <token>         config: token
//	api_key     — custom header or query key             config: key, header (default X-API-Key), in (header|query)
//	basic       — Authorization: Basic base64(u:p)       config: username, password
//	oauth2      — client-credentials token exchange      config: client_id, client_secret, token_url, scopes
//	             (also accepts: token for direct bearer fallback)
//	apple_jwt   — Apple ES256 signed JWT                 config: key_id, issuer_id, private_key, audience
func New(cfg cliSchema.AuthConfig) (Provider, error) {
	resolved := resolveConfig(cfg.Config)

	switch cfg.Type {
	case "apple_jwt":
		return NewAppleJWT(resolved)
	case "bearer":
		// Token may be empty at build time; fail at request time if still missing.
		return &bearerProvider{token: resolved["token"].String()}, nil
	case "api_key":
		key := resolved["key"].String()
		header := resolved["header"].String()
		if header == "" {
			header = "X-API-Key"
		}
		return &apiKeyProvider{key: key, header: header}, nil
	case "basic":
		return &basicProvider{
			username: resolved["username"].String(),
			password: resolved["password"].String(),
		}, nil
	case "oauth2":
		return newOAuth2Provider(resolved)
	case "none", "":
		return &noopProvider{}, nil
	default:
		return nil, fmt.Errorf("unknown auth type: %q (supported: none, bearer, api_key, basic, oauth2, apple_jwt)", cfg.Type)
	}
}

func resolveConfig(cfg map[string]cliSchema.EnvOrValue) map[string]resolvedValue {
	out := make(map[string]resolvedValue, len(cfg))
	for k, ev := range cfg {
		var v string
		switch {
		case ev.Env != "":
			v = os.Getenv(ev.Env)
		case ev.File != "":
			data, err := os.ReadFile(ev.File)
			if err == nil {
				v = string(data)
			}
		default:
			v = ev.Value
		}
		out[k] = resolvedValue{v}
	}
	return out
}

type bearerProvider struct{ token string }

func (b *bearerProvider) AuthHeader() (string, error) {
	if b.token == "" {
		return "", fmt.Errorf("bearer auth: token is not set (configure the 'token' env var in cli-schema auth.config)")
	}
	return "Bearer " + b.token, nil
}

type apiKeyProvider struct{ key, header string }

func (a *apiKeyProvider) AuthHeader() (string, error) {
	return a.key, nil
}

type noopProvider struct{}

func (n *noopProvider) AuthHeader() (string, error) { return "", nil }
