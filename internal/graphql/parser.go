// Package graphql parses GraphQL SDL schema files and generates a CLIForm.
package graphql

import (
	"fmt"
	"os"
	"strings"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/yuraware/clidev/internal/cliSchema"
)

// Spec holds the parsed GraphQL schema.
type Spec struct {
	Title    string
	Version  string
	Endpoint string
	Schema   *ast.Schema
}

// Parse reads a .graphql SDL file and returns a Spec.
func Parse(path string) (*Spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading graphql schema: %w", err)
	}

	schema, gqlErr := gqlparser.LoadSchema(&ast.Source{
		Name:  path,
		Input: string(data),
	})
	if gqlErr != nil {
		return nil, fmt.Errorf("parsing graphql schema: %w", gqlErr)
	}

	return &Spec{
		Title:    "GraphQL API",
		Version:  "1.0",
		Endpoint: "/graphql",
		Schema:   schema,
	}, nil
}

// Generate converts a parsed GraphQL Spec into a CLIForm.
func Generate(spec *Spec) *cliSchema.CLIForm {
	form := &cliSchema.CLIForm{
		Name:        "graphql",
		Version:     spec.Version,
		Description: spec.Title,
		BaseURL:     "https://api.example.com",
		Auth: cliSchema.AuthConfig{
			Type: "bearer",
			Config: map[string]cliSchema.EnvOrValue{
				"token": {Env: "GRAPHQL_TOKEN"},
			},
			Comment: "Set GRAPHQL_TOKEN to your personal access token",
		},
		Output:   cliSchema.OutputConfig{DefaultFormat: "json"},
		Commands: make(map[string]cliSchema.Command),
	}

	if spec.Schema.Query != nil {
		grp := cliSchema.Command{
			Description: "Run GraphQL queries (read operations)",
			Commands:    make(map[string]cliSchema.Command),
		}
		for _, field := range spec.Schema.Query.Fields {
			if strings.HasPrefix(field.Name, "__") {
				continue
			}
			grp.Commands[toKebab(field.Name)] = buildCommand(field, "query", spec)
		}
		if len(grp.Commands) > 0 {
			form.Commands["query"] = grp
		}
	}

	if spec.Schema.Mutation != nil {
		grp := cliSchema.Command{
			Description: "Run GraphQL mutations (write operations)",
			Commands:    make(map[string]cliSchema.Command),
		}
		for _, field := range spec.Schema.Mutation.Fields {
			if strings.HasPrefix(field.Name, "__") {
				continue
			}
			grp.Commands[toKebab(field.Name)] = buildCommand(field, "mutation", spec)
		}
		if len(grp.Commands) > 0 {
			form.Commands["mutate"] = grp
		}
	}

	return form
}

func buildCommand(field *ast.FieldDefinition, opType string, spec *Spec) cliSchema.Command {
	args, bodyFields := buildArgs(field)
	query := buildGraphQLQuery(field, opType, spec)

	cmd := cliSchema.Command{
		Description: field.Description,
		Action: &cliSchema.Action{
			Type:         "graphql",
			Method:       "POST",
			Path:         spec.Endpoint,
			OperationID:  field.Name,
			GraphQLQuery: query,
		},
		Args: args,
	}

	if len(bodyFields) > 0 {
		cmd.Body = &cliSchema.BodyConfig{
			Format:     "graphql_vars",
			Attributes: bodyFields,
		}
	}

	return cmd
}

func buildArgs(field *ast.FieldDefinition) ([]cliSchema.Arg, []cliSchema.BodyField) {
	var args []cliSchema.Arg
	var body []cliSchema.BodyField

	for _, arg := range field.Arguments {
		required := isRequired(arg.Type)
		typeName := gqlTypeToParam(arg.Type)

		if required && isScalarType(arg.Type) {
			args = append(args, cliSchema.Arg{
				Name:     arg.Name,
				Required: true,
				Desc:     arg.Description,
			})
		} else {
			body = append(body, cliSchema.BodyField{
				Flag:     "--" + toKebab(arg.Name),
				Field:    arg.Name,
				Type:     typeName,
				Required: required,
				Desc:     arg.Description,
			})
		}
	}
	return args, body
}

// buildGraphQLQuery generates a minimal GraphQL operation string.
func buildGraphQLQuery(field *ast.FieldDefinition, opType string, spec *Spec) string {
	var vars, varPass []string
	for _, arg := range field.Arguments {
		gqlType := renderGQLType(arg.Type)
		vars = append(vars, fmt.Sprintf("$%s: %s", arg.Name, gqlType))
		varPass = append(varPass, fmt.Sprintf("%s: $%s", arg.Name, arg.Name))
	}

	// Select scalar fields from the return type.
	selection := buildSelection(field.Type, spec.Schema, 0)

	varDecl := ""
	if len(vars) > 0 {
		varDecl = "(" + strings.Join(vars, ", ") + ")"
	}
	argCall := ""
	if len(varPass) > 0 {
		argCall = "(" + strings.Join(varPass, ", ") + ")"
	}

	return fmt.Sprintf("%s %s%s { %s%s { %s } }",
		opType, field.Name, varDecl, field.Name, argCall, selection)
}

// buildSelection picks scalar fields from a type up to depth 1.
func buildSelection(t *ast.Type, schema *ast.Schema, depth int) string {
	if depth > 1 {
		return "__typename"
	}
	typeName := t.NamedType
	if typeName == "" && t.Elem != nil {
		typeName = t.Elem.NamedType
	}

	def, ok := schema.Types[typeName]
	if !ok || def.Kind != ast.Object {
		return "__typename"
	}

	var fields []string
	for _, f := range def.Fields {
		if strings.HasPrefix(f.Name, "__") {
			continue
		}
		if isLeafType(f.Type, schema) {
			fields = append(fields, f.Name)
		}
		if len(fields) >= 10 {
			break
		}
	}
	if len(fields) == 0 {
		return "__typename"
	}
	return strings.Join(fields, " ")
}

func isLeafType(t *ast.Type, schema *ast.Schema) bool {
	name := t.NamedType
	if name == "" && t.Elem != nil {
		name = t.Elem.NamedType
	}
	def, ok := schema.Types[name]
	if !ok {
		return false
	}
	return def.Kind == ast.Scalar || def.Kind == ast.Enum
}

func isRequired(t *ast.Type) bool {
	return t.NonNull
}

func isScalarType(t *ast.Type) bool {
	name := t.NamedType
	if name == "" {
		return false
	}
	switch name {
	case "String", "Int", "Float", "Boolean", "ID":
		return true
	}
	return false
}

func gqlTypeToParam(t *ast.Type) string {
	name := t.NamedType
	if name == "" && t.Elem != nil {
		return "string_array"
	}
	switch name {
	case "Int", "Float":
		return "integer"
	case "Boolean":
		return "boolean"
	default:
		return "string"
	}
}

func renderGQLType(t *ast.Type) string {
	if t.Elem != nil {
		inner := renderGQLType(t.Elem)
		if t.NonNull {
			return "[" + inner + "]!"
		}
		return "[" + inner + "]"
	}
	if t.NonNull {
		return t.NamedType + "!"
	}
	return t.NamedType
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
