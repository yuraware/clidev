package cliSchema

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Load reads a cli-schema file (YAML or JSON) and returns a CLIForm.
func Load(path string) (*CLIForm, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading cli-schema: %w", err)
	}

	var form CLIForm
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &form); err != nil {
			return nil, fmt.Errorf("parsing cli-schema JSON: %w", err)
		}
	default: // .yaml, .yml, or anything else — try YAML
		if err := yaml.Unmarshal(data, &form); err != nil {
			return nil, fmt.Errorf("parsing cli-schema YAML: %w", err)
		}
	}

	if err := validate(&form); err != nil {
		return nil, err
	}
	return &form, nil
}

func validate(f *CLIForm) error {
	if f.Name == "" {
		return fmt.Errorf("cli-schema: missing required field 'name'")
	}
	if f.BaseURL == "" {
		return fmt.Errorf("cli-schema: missing required field 'base_url'")
	}
	return nil
}
