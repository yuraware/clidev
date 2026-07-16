package cliSchema

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// JSONSchemaRoot is the JSON Schema that validates a cli-schema document.
var JSONSchemaRoot = map[string]any{
	"$schema":     "https://json-schema.org/draft/2020-12/schema",
	"$id":         "https://github.com/yuraware/clidev/schema/cli-schema.schema.json",
	"title":       "cli-schema",
	"description": "Declarative CLI definition format for clidev",
	"type":        "object",
	"required":    []string{"name", "base_url"},
	"properties": map[string]any{
		"name":        map[string]any{"type": "string", "description": "CLI tool name (used as the root command)"},
		"version":     map[string]any{"type": "string"},
		"description": map[string]any{"type": "string"},
		"base_url":    map[string]any{"type": "string", "format": "uri"},
		"auth": map[string]any{
			"type": "object",
			"properties": map[string]any{
				"type": map[string]any{
					"type": "string",
					"enum": []string{"none", "bearer", "api_key", "apple_jwt"},
				},
				"config": map[string]any{
					"type":                 "object",
					"additionalProperties": envOrValueSchema(),
				},
			},
		},
		"output": map[string]any{
			"type": "object",
			"properties": map[string]any{
				"default_format": map[string]any{
					"type": "string",
					"enum": []string{"table", "json", "yaml"},
				},
			},
		},
		"commands": map[string]any{
			"type":                 "object",
			"additionalProperties": commandSchema(),
		},
	},
}

func envOrValueSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"env":   map[string]any{"type": "string"},
			"value": map[string]any{"type": "string"},
			"file":  map[string]any{"type": "string"},
		},
	}
}

func commandSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"description": map[string]any{"type": "string"},
			"action": map[string]any{
				"type":     "object",
				"required": []string{"method", "path"},
				"properties": map[string]any{
					"method":       map[string]any{"type": "string", "enum": []string{"GET", "POST", "PATCH", "DELETE", "PUT"}},
					"path":         map[string]any{"type": "string"},
					"operation_id": map[string]any{"type": "string"},
				},
			},
			"args": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type":     "object",
					"required": []string{"name"},
					"properties": map[string]any{
						"name":        map[string]any{"type": "string"},
						"path_param":  map[string]any{"type": "string"},
						"required":    map[string]any{"type": "boolean"},
						"description": map[string]any{"type": "string"},
					},
				},
			},
			"parameters": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type":     "object",
					"required": []string{"flag", "query", "type"},
					"properties": map[string]any{
						"flag":        map[string]any{"type": "string"},
						"query":       map[string]any{"type": "string"},
						"type":        map[string]any{"type": "string", "enum": []string{"string", "string_array", "integer", "boolean", "enum", "enum_array"}},
						"values":      map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
						"default":     map[string]any{},
						"max":         map[string]any{"type": "integer"},
						"description": map[string]any{"type": "string"},
					},
				},
			},
			"body": map[string]any{
				"type":     "object",
				"required": []string{"format"},
				"properties": map[string]any{
					"format":        map[string]any{"type": "string", "enum": []string{"json_api", "raw_json"}},
					"resource_type": map[string]any{"type": "string"},
					"attributes": map[string]any{
						"type": "array",
						"items": map[string]any{
							"type":     "object",
							"required": []string{"flag", "field", "type"},
							"properties": map[string]any{
								"flag":        map[string]any{"type": "string"},
								"field":       map[string]any{"type": "string"},
								"type":        map[string]any{"type": "string"},
								"required":    map[string]any{"type": "boolean"},
								"default":     map[string]any{},
								"description": map[string]any{"type": "string"},
							},
						},
					},
					"relationships": map[string]any{
						"type": "array",
						"items": map[string]any{
							"type":     "object",
							"required": []string{"flag", "relationship", "resource_type"},
							"properties": map[string]any{
								"flag":          map[string]any{"type": "string"},
								"relationship":  map[string]any{"type": "string"},
								"resource_type": map[string]any{"type": "string"},
								"required":      map[string]any{"type": "boolean"},
							},
						},
					},
				},
			},
			"commands": map[string]any{
				"type":                 "object",
				"additionalProperties": map[string]any{"$ref": "#/properties/commands/additionalProperties"},
			},
		},
	}
}

// WriteSchema writes the JSON Schema to path (JSON or YAML based on extension).
func WriteSchema(path string) error {
	ext := strings.ToLower(filepath.Ext(path))
	var data []byte
	var err error

	switch ext {
	case ".json":
		data, err = json.MarshalIndent(JSONSchemaRoot, "", "  ")
		if err != nil {
			return fmt.Errorf("marshalling schema JSON: %w", err)
		}
		data = append(data, '\n')
	default:
		data, err = yaml.Marshal(JSONSchemaRoot)
		if err != nil {
			return fmt.Errorf("marshalling schema YAML: %w", err)
		}
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("writing schema: %w", err)
	}
	return nil
}
