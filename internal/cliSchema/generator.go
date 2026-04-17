package cliSchema

import (
	"fmt"
	"strings"

	"github.com/yurikobets/cli-builder/internal/oas"
)

// Generate converts a parsed OAS spec into a CLIForm.
func Generate(spec *oas.Spec, ops []oas.Op) *CLIForm {
	baseURL := ""
	if len(spec.Servers) > 0 {
		baseURL = strings.TrimRight(spec.Servers[0].URL, "/")
	}

	form := &CLIForm{
		Name:        slugify(spec.Info.Title),
		Version:     spec.Info.Version,
		Description: spec.Info.Title,
		BaseURL:     baseURL,
		Auth: AuthConfig{
			Type: "none",
		},
		Output: OutputConfig{
			DefaultFormat: "table",
		},
		Commands: make(map[string]Command),
	}

	for _, op := range ops {
		placeOp(form, op)
	}

	return form
}

// placeOp maps a single operation into the command tree.
func placeOp(form *CLIForm, op oas.Op) {
	parts := parseOperationID(op.OperationID)
	if parts == nil {
		return
	}

	resource := parts[0]
	action := parts[1]
	sub := ""
	if len(parts) == 3 {
		sub = parts[1]
		action = parts[2]
	}

	// Map action names to CLI verbs.
	verb := actionToVerb(action)
	if verb == "" {
		return // skip relationship metadata operations
	}

	// Ensure the resource group exists.
	grp, ok := form.Commands[resource]
	if !ok {
		grp = Command{
			Description: fmt.Sprintf("Manage %s", resource),
			Commands:    make(map[string]Command),
		}
	}
	if grp.Commands == nil {
		grp.Commands = make(map[string]Command)
	}

	cmd := buildCommand(op, verb)

	if sub != "" {
		// Nested: resource → sub → verb
		subGrp, ok := grp.Commands[sub]
		if !ok {
			subGrp = Command{
				Description: fmt.Sprintf("Manage %s %s", resource, sub),
				Commands:    make(map[string]Command),
			}
		}
		if subGrp.Commands == nil {
			subGrp.Commands = make(map[string]Command)
		}
		subGrp.Commands[verb] = cmd
		grp.Commands[sub] = subGrp
	} else {
		grp.Commands[verb] = cmd
	}

	form.Commands[resource] = grp
}

// buildCommand creates a leaf Command for one operation.
func buildCommand(op oas.Op, verb string) Command {
	cmd := Command{
		Description: fmt.Sprintf("%s %s", strings.Title(verb), op.Path),
		Action: &Action{
			Method:      op.Method,
			Path:        op.Path,
			OperationID: op.OperationID,
		},
	}

	// Extract path parameters as positional args.
	for _, seg := range strings.Split(op.Path, "/") {
		if strings.HasPrefix(seg, "{") && strings.HasSuffix(seg, "}") {
			name := seg[1 : len(seg)-1]
			cmd.Args = append(cmd.Args, Arg{
				Name:      name,
				PathParam: name,
				Required:  true,
				Desc:      fmt.Sprintf("ID of the %s resource", name),
			})
		}
	}

	// Map query parameters to flags.
	for _, p := range op.Parameters {
		if p.In != "query" {
			continue
		}
		param := queryParamToParameter(p)
		if param != nil {
			cmd.Parameters = append(cmd.Parameters, *param)
		}
	}

	// Mark body if present.
	if op.HasBody {
		cmd.Body = &BodyConfig{
			Format:       "json_api",
			ResourceType: extractResourceType(op.Path),
		}
	}

	return cmd
}

// queryParamToParameter converts an OAS query parameter to a cli-schema Parameter.
func queryParamToParameter(p oas.Parameter) *Parameter {
	name := p.Name
	flag := "--" + paramNameToFlag(name)

	paramType := inferParamType(name, p.Schema)

	param := &Parameter{
		Flag:  flag,
		Query: name,
		Type:  paramType,
		Desc:  p.Description,
	}

	// Extract enum values from schema.
	if p.Schema != nil {
		schema := p.Schema
		// For array params the enum lives in items.
		if schema.Items != nil {
			schema = schema.Items
		}
		for _, v := range schema.Enum {
			param.Values = append(param.Values, fmt.Sprintf("%v", v))
		}
		if schema.Maximum != nil {
			param.Max = int(*schema.Maximum)
		}
	}

	return param
}

// inferParamType derives a cli-schema type string from the param name and schema.
func inferParamType(name string, schema *oas.SchemaRef) string {
	// Bracket-notation arrays: fields[x], filter[x], limit[x], include
	if strings.Contains(name, "[") || name == "include" || name == "sort" {
		if schema != nil && schema.Type == "array" {
			if schema.Items != nil && len(schema.Items.Enum) > 0 {
				return "enum_array"
			}
			return "string_array"
		}
		// sort is a single enum
		if name == "sort" {
			return "enum"
		}
		return "string_array"
	}

	if schema == nil {
		return "string"
	}

	switch schema.Type {
	case "integer":
		return "integer"
	case "boolean":
		return "boolean"
	case "array":
		if schema.Items != nil && len(schema.Items.Enum) > 0 {
			return "enum_array"
		}
		return "string_array"
	default:
		if len(schema.Enum) > 0 {
			return "enum"
		}
		return "string"
	}
}

// paramNameToFlag converts OAS param names like `filter[bundleId]` to `filter-bundle-id`.
func paramNameToFlag(name string) string {
	// fields[apps] → fields-apps
	name = strings.ReplaceAll(name, "[", "-")
	name = strings.ReplaceAll(name, "]", "")
	// camelCase → kebab-case
	var b strings.Builder
	for i, r := range name {
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

// parseOperationID splits `apps_events_getCollection` into ["apps","events","getCollection"].
// Returns nil for formats we don't handle (relationship metadata, etc.).
func parseOperationID(id string) []string {
	parts := strings.SplitN(id, "_", 3)
	if len(parts) < 2 {
		return nil
	}
	// Normalise to [resource, (sub,) action]
	last := parts[len(parts)-1]
	if !isKnownAction(last) {
		return nil
	}
	switch len(parts) {
	case 2:
		return []string{parts[0], parts[1]}
	case 3:
		return []string{parts[0], parts[1], parts[2]}
	}
	return nil
}

var knownActions = map[string]bool{
	"getCollection":           true,
	"getInstance":             true,
	"createInstance":          true,
	"updateInstance":          true,
	"deleteInstance":          true,
	"getToManyRelated":        true,
	"getToOneRelated":         true,
	"getToManyRelationship":   true,
	"getToOneRelationship":    true,
	"createToManyRelationship": true,
	"deleteToManyRelationship": true,
	"replaceToManyRelationship": true,
	"updateToOneRelationship":  true,
	"getMetrics":              true,
}

func isKnownAction(s string) bool {
	return knownActions[s]
}

// actionToVerb maps OAS action names to CLI verbs.
// Returns "" for relationship metadata operations we skip.
func actionToVerb(action string) string {
	switch action {
	case "getCollection", "getToManyRelated":
		return "list"
	case "getInstance", "getToOneRelated":
		return "get"
	case "createInstance":
		return "create"
	case "updateInstance":
		return "update"
	case "deleteInstance":
		return "delete"
	case "getMetrics":
		return "metrics"
	default:
		// Skip relationship metadata and relationship mutation operations.
		return ""
	}
}

func extractResourceType(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	// /v1/apps → apps; /v1/apps/{id}/events → apps
	if len(parts) >= 2 {
		return parts[1]
	}
	return "unknown"
}

func slugify(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
}
