package auth

import (
	"fmt"
	"os"

	"github.com/yurikobets/cli-builder/internal/cliSchema"
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
func New(cfg cliSchema.AuthConfig) (Provider, error) {
	resolved := resolveConfig(cfg.Config)

	switch cfg.Type {
	case "apple_jwt":
		return NewAppleJWT(resolved)
	case "bearer":
		token := resolved["token"].String()
		if token == "" {
			return nil, fmt.Errorf("bearer auth: missing 'token' config key")
		}
		return &bearerProvider{token: token}, nil
	case "api_key":
		key := resolved["key"].String()
		header := resolved["header"].String()
		if header == "" {
			header = "X-API-Key"
		}
		return &apiKeyProvider{key: key, header: header}, nil
	case "none", "":
		return &noopProvider{}, nil
	default:
		return nil, fmt.Errorf("unknown auth type: %q", cfg.Type)
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
	return "Bearer " + b.token, nil
}

type apiKeyProvider struct{ key, header string }

func (a *apiKeyProvider) AuthHeader() (string, error) {
	return a.key, nil // caller uses the custom header name
}

type noopProvider struct{}

func (n *noopProvider) AuthHeader() (string, error) { return "", nil }
