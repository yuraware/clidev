package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/yurikobets/cli-builder/internal/cliSchema"
	"github.com/yurikobets/cli-builder/internal/oas"
)

func main() {
	root := &cobra.Command{
		Use:   "builder",
		Short: "Generate cli-schema files from API specs",
	}

	root.AddCommand(generateCmd())
	root.AddCommand(schemaCmd())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func generateCmd() *cobra.Command {
	var specPath, outPath string

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a cli-schema file from an OpenAPI spec",
		Example: `  builder generate --spec sample-api/acs/openapi.oas.json --out sample-api/acs/cli-schema.yaml
  builder generate --spec api.json --out cli-schema.json`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if specPath == "" {
				return fmt.Errorf("--spec is required")
			}

			fmt.Fprintf(os.Stderr, "Parsing %s...\n", specPath)
			spec, ops, err := oas.Parse(specPath)
			if err != nil {
				return fmt.Errorf("parsing spec: %w", err)
			}
			fmt.Fprintf(os.Stderr, "Found %d operations across %d paths\n", len(ops), len(spec.Paths))

			form := cliSchema.Generate(spec, ops)

			if outPath == "" {
				outPath = "cli-schema.yaml"
			}

			if err := writeForm(form, outPath); err != nil {
				return err
			}

			fmt.Fprintf(os.Stderr, "Written: %s\n", outPath)

			// Also emit matching schema file.
			schemaPath := schemaPathFor(outPath)
			if err := cliSchema.WriteSchema(schemaPath); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: could not write schema: %v\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "Schema:  %s\n", schemaPath)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&specPath, "spec", "", "Path to OpenAPI spec (JSON or YAML)")
	cmd.Flags().StringVar(&outPath, "out", "", "Output cli-schema path (default: cli-schema.yaml)")
	return cmd
}

func schemaCmd() *cobra.Command {
	var outPath string

	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Write the cli-schema JSON Schema to a file",
		Example: `  builder schema --out schema/cli-schema.schema.json
  builder schema --out schema/cli-schema.schema.yaml`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if outPath == "" {
				outPath = "schema/cli-schema.schema.json"
			}
			if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
				return err
			}
			if err := cliSchema.WriteSchema(outPath); err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "Written: %s\n", outPath)
			return nil
		},
	}

	cmd.Flags().StringVar(&outPath, "out", "", "Output path (default: schema/cli-schema.schema.json)")
	return cmd
}

func writeForm(form *cliSchema.CLIForm, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	ext := strings.ToLower(filepath.Ext(path))
	var data []byte
	var err error

	switch ext {
	case ".json":
		data, err = json.MarshalIndent(form, "", "  ")
		if err != nil {
			return fmt.Errorf("marshalling JSON: %w", err)
		}
		data = append(data, '\n')
	default:
		data, err = yaml.Marshal(form)
		if err != nil {
			return fmt.Errorf("marshalling YAML: %w", err)
		}
	}

	return os.WriteFile(path, data, 0644)
}

// schemaPathFor returns a sibling schema file path for a given cli-schema path.
func schemaPathFor(cliFormPath string) string {
	dir := filepath.Dir(cliFormPath)
	base := strings.TrimSuffix(filepath.Base(cliFormPath), filepath.Ext(cliFormPath))
	ext := filepath.Ext(cliFormPath)
	return filepath.Join(dir, base+".schema"+ext)
}
