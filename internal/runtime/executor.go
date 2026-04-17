package runtime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yurikobets/cli-builder/internal/cliSchema"
	"github.com/yurikobets/cli-builder/internal/runtime/auth"
)

// Executor runs a CLI command's HTTP action.
type Executor struct {
	BaseURL  string
	Auth     auth.Provider
	client   *http.Client
}

// NewExecutor creates a ready-to-use Executor.
func NewExecutor(baseURL string, authProvider auth.Provider) *Executor {
	return &Executor{
		BaseURL: strings.TrimRight(baseURL, "/"),
		Auth:    authProvider,
		client:  &http.Client{},
	}
}

// Execute resolves flags/args from a Cobra command and calls the API.
func (e *Executor) Execute(cobraCmd *cobra.Command, args []string, cmd cliSchema.Command) ([]byte, error) {
	if cmd.Action == nil {
		return nil, fmt.Errorf("command has no action")
	}
	action := cmd.Action

	// 1. Build path: substitute {param} placeholders with positional args.
	path := action.Path
	for i, arg := range cmd.Args {
		if i >= len(args) {
			if arg.Required {
				return nil, fmt.Errorf("missing required argument: %s", arg.Name)
			}
			break
		}
		placeholder := "{" + arg.PathParam + "}"
		path = strings.ReplaceAll(path, placeholder, args[i])
	}

	// 2. Build query string from flags.
	queryParams := url.Values{}
	for _, p := range cmd.Parameters {
		flagName := strings.TrimPrefix(p.Flag, "--")
		val, err := getFlag(cobraCmd, flagName, p.Type)
		if err != nil || val == "" {
			continue
		}
		// Multi-value flags (string_array / enum_array) → comma-separated or repeated.
		if strings.HasSuffix(p.Type, "_array") {
			for _, item := range strings.Split(val, ",") {
				queryParams.Add(p.Query, strings.TrimSpace(item))
			}
		} else {
			queryParams.Set(p.Query, val)
		}
	}

	rawURL := e.BaseURL + path
	if len(queryParams) > 0 {
		rawURL += "?" + queryParams.Encode()
	}

	// 3. Build request body.
	var bodyReader io.Reader
	if cmd.Body != nil {
		bodyBytes, err := buildBody(cobraCmd, cmd.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(action.Method, rawURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}

	if cmd.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	// 4. Attach auth header.
	authHeader, err := e.Auth.AuthHeader()
	if err != nil {
		return nil, fmt.Errorf("generating auth: %w", err)
	}
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	// 5. Execute.
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// getFlag reads a flag value as a string regardless of underlying type.
func getFlag(cmd *cobra.Command, name, paramType string) (string, error) {
	f := cmd.Flags().Lookup(name)
	if f == nil || f.Value.String() == f.DefValue {
		return "", nil // not set
	}
	return f.Value.String(), nil
}

// buildBody assembles a JSON request body from CLI flags.
func buildBody(cmd *cobra.Command, body *cliSchema.BodyConfig) ([]byte, error) {
	switch body.Format {
	case "json_api":
		return buildJSONAPIBody(cmd, body)
	default:
		return buildRawBody(cmd, body)
	}
}

func buildJSONAPIBody(cmd *cobra.Command, body *cliSchema.BodyConfig) ([]byte, error) {
	attrs := map[string]any{}
	for _, f := range body.Attributes {
		flagName := strings.TrimPrefix(f.Flag, "--")
		fl := cmd.Flags().Lookup(flagName)
		if fl == nil {
			continue
		}
		val := fl.Value.String()
		if val == "" && f.Required {
			return nil, fmt.Errorf("missing required flag: %s", f.Flag)
		}
		if val != "" {
			attrs[f.Field] = val
		} else if f.Default != nil {
			attrs[f.Field] = f.Default
		}
	}

	rels := map[string]any{}
	for _, r := range body.Relationships {
		flagName := strings.TrimPrefix(r.Flag, "--")
		fl := cmd.Flags().Lookup(flagName)
		if fl == nil {
			continue
		}
		val := fl.Value.String()
		if val == "" {
			if r.Required {
				return nil, fmt.Errorf("missing required relationship flag: %s", r.Flag)
			}
			continue
		}
		rels[r.Relationship] = map[string]any{
			"data": map[string]any{"type": r.ResourceType, "id": val},
		}
	}

	data := map[string]any{
		"type":       body.ResourceType,
		"attributes": attrs,
	}
	if len(rels) > 0 {
		data["relationships"] = rels
	}

	return json.Marshal(map[string]any{"data": data})
}

func buildRawBody(cmd *cobra.Command, body *cliSchema.BodyConfig) ([]byte, error) {
	m := map[string]any{}
	for _, f := range body.Attributes {
		flagName := strings.TrimPrefix(f.Flag, "--")
		fl := cmd.Flags().Lookup(flagName)
		if fl == nil {
			continue
		}
		val := fl.Value.String()
		if val != "" {
			m[f.Field] = val
		}
	}
	return json.Marshal(m)
}
