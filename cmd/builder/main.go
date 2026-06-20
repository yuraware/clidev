package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/yurikobets/cli-builder/internal/asyncapi"
	"github.com/yurikobets/cli-builder/internal/cliSchema"
	"github.com/yurikobets/cli-builder/internal/formats"
	gql "github.com/yurikobets/cli-builder/internal/graphql"
	"github.com/yurikobets/cli-builder/internal/oas"
	"github.com/yurikobets/cli-builder/internal/proto"
)

func main() {
	root := &cobra.Command{
		Use:   "builder",
		Short: "Generate cli-schema files from API specs",
		Long: `builder converts API specs in multiple formats into a declarative cli-schema file.

Supported input formats:
  openapi   — OpenAPI 3.x / Swagger 2.x  (.json, .yaml, .yml)
  proto     — Protocol Buffers gRPC       (.proto)
  graphql   — GraphQL SDL schema          (.graphql, .gql)
  asyncapi  — AsyncAPI 2.x / 3.x         (.yaml, .json)`,
	}

	root.AddCommand(generateCmd())
	root.AddCommand(schemaCmd())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func generateCmd() *cobra.Command {
	var specPath, outPath, forceFmt string

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a cli-schema from an API spec (auto-detects format)",
		Example: `  # OpenAPI
  builder generate --spec sample-api/acs/openapi.oas.json --out sample-api/acs/cli-schema.yaml

  # gRPC / Protocol Buffers
  builder generate --spec sample-api/pubsub/pubsub.proto --out sample-api/pubsub/cli-schema.yaml

  # GraphQL SDL
  builder generate --spec sample-api/github-gql/schema.graphql --out sample-api/github-gql/cli-schema.yaml

  # AsyncAPI
  builder generate --spec sample-api/streetlights/asyncapi.yml --out sample-api/streetlights/cli-schema.yaml

  # Force a specific format
  builder generate --spec api.yaml --format openapi --out cli-schema.yaml`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if specPath == "" {
				return fmt.Errorf("--spec is required")
			}

			format := forceFmt
			if format == "" {
				format = formats.Detect(specPath)
			}
			if format == formats.Unknown {
				return fmt.Errorf("cannot detect format for %q — use --format (openapi|proto|graphql|asyncapi)", specPath)
			}

			fmt.Fprintf(os.Stderr, "Format:  %s\n", format)
			fmt.Fprintf(os.Stderr, "Parsing %s...\n", specPath)

			form, stats, err := generateForm(specPath, format)
			if err != nil {
				return err
			}

			if outPath == "" {
				outPath = "cli-schema.yaml"
			}
			if err := writeForm(form, outPath); err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "Written: %s\n", outPath)

			schemaPath := schemaPathFor(outPath)
			if err := cliSchema.WriteSchema(schemaPath); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: could not write schema: %v\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "Schema:  %s\n", schemaPath)
			}

			docPath := cliSchema.DocPathFor(outPath)
			docOpts := cliSchema.DocOptions{
				SpecPath:     specPath,
				SchemaPath:   outPath,
				Format:       format,
				Stats:        stats,
				FormatForced: forceFmt != "",
			}
			if err := cliSchema.WriteDoc(docPath, form, docOpts); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: could not write docs: %v\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "Docs:    %s\n", docPath)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&specPath, "spec", "", "Path to API spec file")
	cmd.Flags().StringVar(&outPath, "out", "", "Output cli-schema path (default: cli-schema.yaml)")
	cmd.Flags().StringVar(&forceFmt, "format", "", "Force input format: openapi | proto | graphql | asyncapi")
	return cmd
}

// generateForm parses the spec and returns a CLIForm and human-readable stats for the detected format.
func generateForm(specPath, format string) (*cliSchema.CLIForm, string, error) {
	switch format {
	case formats.OpenAPI:
		spec, ops, err := oas.Parse(specPath)
		if err != nil {
			return nil, "", fmt.Errorf("parsing OpenAPI: %w", err)
		}
		fmt.Fprintf(os.Stderr, "Found %d operations across %d paths\n", len(ops), len(spec.Paths))
		stats := fmt.Sprintf("%d paths · %d operations", len(spec.Paths), len(ops))
		return cliSchema.Generate(spec, ops), stats, nil

	case formats.Proto:
		spec, err := proto.Parse(specPath)
		if err != nil {
			return nil, "", fmt.Errorf("parsing proto: %w", err)
		}
		total := 0
		for _, s := range spec.Services {
			total += len(s.RPCs)
		}
		fmt.Fprintf(os.Stderr, "Found %d services, %d RPCs\n", len(spec.Services), total)
		stats := fmt.Sprintf("%d services · %d RPCs", len(spec.Services), total)
		return proto.Generate(spec), stats, nil

	case formats.GraphQL:
		spec, err := gql.Parse(specPath)
		if err != nil {
			return nil, "", fmt.Errorf("parsing GraphQL: %w", err)
		}
		qCount, mCount := 0, 0
		if spec.Schema.Query != nil {
			qCount = len(spec.Schema.Query.Fields)
		}
		if spec.Schema.Mutation != nil {
			mCount = len(spec.Schema.Mutation.Fields)
		}
		fmt.Fprintf(os.Stderr, "Found %d queries, %d mutations\n", qCount, mCount)
		stats := fmt.Sprintf("%d queries · %d mutations", qCount, mCount)
		return gql.Generate(spec), stats, nil

	case formats.AsyncAPI:
		spec, err := asyncapi.Parse(specPath)
		if err != nil {
			return nil, "", fmt.Errorf("parsing AsyncAPI: %w", err)
		}
		fmt.Fprintf(os.Stderr, "AsyncAPI %s — %d channels, %d operations\n",
			spec.AsyncAPI, len(spec.Channels), len(spec.Operations))
		stats := fmt.Sprintf("AsyncAPI %s · %d channels · %d operations",
			spec.AsyncAPI, len(spec.Channels), len(spec.Operations))
		return asyncapi.Generate(spec), stats, nil

	default:
		return nil, "", fmt.Errorf("unsupported format: %q", format)
	}
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

func schemaPathFor(p string) string {
	dir := filepath.Dir(p)
	base := strings.TrimSuffix(filepath.Base(p), filepath.Ext(p))
	ext := filepath.Ext(p)
	return filepath.Join(dir, base+".schema"+ext)
}
