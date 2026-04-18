package proto

import (
	"fmt"
	"strings"

	"github.com/yurikobets/cli-builder/internal/cliSchema"
)

// Generate converts a parsed proto Spec into a CLIForm.
func Generate(spec *Spec) *cliSchema.CLIForm {
	baseURL := ""
	if spec.DefaultHost != "" {
		baseURL = "https://" + spec.DefaultHost
	}

	name := slugify(spec.PackageName)

	form := &cliSchema.CLIForm{
		Name:        name,
		Version:     "1.0",
		Description: fmt.Sprintf("%s gRPC API (HTTP transcoding)", spec.PackageName),
		BaseURL:     baseURL,
		Auth:        inferAuth(spec),
		Output:      cliSchema.OutputConfig{DefaultFormat: "table"},
		Commands:    make(map[string]cliSchema.Command),
	}

	for _, svc := range spec.Services {
		grpName := toKebab(svc.Name)
		grp := cliSchema.Command{
			Description: fmt.Sprintf("%s service", svc.Name),
			Commands:    make(map[string]cliSchema.Command),
		}

		for _, rpc := range svc.RPCs {
			if rpc.HTTPMethod == "" || rpc.HTTPPath == "" {
				continue
			}
			verb := rpcToVerb(rpc.Name)
			cmd := buildCommand(rpc, spec)
			// Avoid overwriting: append service suffix if collision.
			key := verb
			if _, exists := grp.Commands[key]; exists {
				key = verb + "-" + toKebab(rpc.Name)
			}
			grp.Commands[key] = cmd
		}

		if len(grp.Commands) > 0 {
			form.Commands[grpName] = grp
		}
	}

	return form
}

func buildCommand(rpc RPC, spec *Spec) cliSchema.Command {
	cmd := cliSchema.Command{
		Description: fmt.Sprintf("%s → %s %s", rpc.Name, rpc.HTTPMethod, rpc.HTTPPath),
		Action: &cliSchema.Action{
			Method:      rpc.HTTPMethod,
			Path:        rpc.HTTPPath,
			OperationID: rpc.Name,
		},
	}

	// Path parameters as positional args.
	for _, seg := range strings.Split(rpc.HTTPPath, "/") {
		if strings.HasPrefix(seg, "{") && strings.HasSuffix(seg, "}") {
			name := seg[1 : len(seg)-1]
			cmd.Args = append(cmd.Args, cliSchema.Arg{
				Name:      name,
				PathParam: name,
				Required:  true,
				Desc:      fmt.Sprintf("Resource name / %s", name),
			})
		}
	}

	// Input message fields → flags (for POST/PUT/PATCH with body).
	if (rpc.HTTPMethod == "POST" || rpc.HTTPMethod == "PUT" || rpc.HTTPMethod == "PATCH") && rpc.Body != "" {
		msg, ok := spec.Messages[rpc.InputType]
		if ok {
			body := &cliSchema.BodyConfig{Format: "raw_json"}
			for _, f := range msg.Fields {
				if isPathParam(f.Name, rpc.HTTPPath) {
					continue
				}
				body.Attributes = append(body.Attributes, cliSchema.BodyField{
					Flag:  "--" + toKebab(f.Name),
					Field: f.Name,
					Type:  protoTypeToParamType(f.Type, f.Repeated),
				})
			}
			if len(body.Attributes) > 0 {
				cmd.Body = body
			}
		}
	}

	return cmd
}

func isPathParam(field, path string) bool {
	return strings.Contains(path, "{"+field+"}")
}

func inferAuth(spec *Spec) cliSchema.AuthConfig {
	if len(spec.OAuthScopes) > 0 {
		scope := ""
		if len(spec.OAuthScopes) > 0 {
			scope = spec.OAuthScopes[0]
		}
		return cliSchema.AuthConfig{
			Type: "bearer",
			Config: map[string]cliSchema.EnvOrValue{
				"token": {Env: "GOOGLE_ACCESS_TOKEN"},
			},
			Comment: fmt.Sprintf("OAuth2 scopes: %s — obtain a token via: gcloud auth print-access-token", scope),
		}
	}
	return cliSchema.AuthConfig{Type: "none"}
}

func rpcToVerb(name string) string {
	lower := strings.ToLower(name)
	switch {
	case strings.HasPrefix(lower, "list") || strings.HasPrefix(lower, "search"):
		return "list"
	case strings.HasPrefix(lower, "get"):
		return "get"
	case strings.HasPrefix(lower, "create"):
		return "create"
	case strings.HasPrefix(lower, "update") || strings.HasPrefix(lower, "patch"):
		return "update"
	case strings.HasPrefix(lower, "delete"):
		return "delete"
	case strings.HasPrefix(lower, "publish"):
		return "publish"
	case strings.HasPrefix(lower, "pull"):
		return "pull"
	case strings.HasPrefix(lower, "acknowledge"):
		return "acknowledge"
	default:
		return toKebab(name)
	}
}

func protoTypeToParamType(t string, repeated bool) string {
	if repeated {
		return "string_array"
	}
	switch t {
	case "int32", "int64", "uint32", "uint64", "fixed32", "fixed64", "sfixed32", "sfixed64":
		return "integer"
	case "bool":
		return "boolean"
	default:
		return "string"
	}
}

func toKebab(s string) string {
	var b strings.Builder
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				b.WriteRune('-')
			}
			b.WriteRune(r + 32)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, ".", "-")
	s = strings.ReplaceAll(s, "_", "-")
	return s
}
