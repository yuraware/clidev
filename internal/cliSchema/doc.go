package cliSchema

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const maxFullCommandDocs = 100
const maxExampleCommands = 10

// DocOptions supplies context for generated CLI documentation.
type DocOptions struct {
	SpecPath     string
	SchemaPath   string
	Format       string
	Stats        string
	FormatForced bool
}

// DocPathFor returns the markdown doc path alongside a cli-schema output file.
func DocPathFor(schemaPath string) string {
	dir := filepath.Dir(schemaPath)
	base := strings.TrimSuffix(filepath.Base(schemaPath), filepath.Ext(schemaPath))
	return filepath.Join(dir, base+".md")
}

// WriteDoc writes a markdown description file for a generated CLI schema.
func WriteDoc(path string, form *CLIForm, opts DocOptions) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	content := GenerateMarkdown(form, opts)
	return os.WriteFile(path, []byte(content), 0644)
}

// GenerateMarkdown builds the markdown documentation body.
func GenerateMarkdown(form *CLIForm, opts DocOptions) string {
	var b strings.Builder

	title := form.Description
	if title == "" {
		title = form.Name
	}
	b.WriteString("# " + title + "\n\n")

	if form.Description != "" && form.Description != title {
		b.WriteString(form.Description + "\n\n")
	}

	b.WriteString("**CLI name:** `" + form.Name + "`")
	if form.Version != "" {
		b.WriteString(" · **Version:** " + form.Version)
	}
	b.WriteString("\n")

	if opts.Format != "" {
		b.WriteString("**Spec format:** " + formatLabel(opts.Format))
		if opts.Stats != "" {
			b.WriteString(" · " + opts.Stats)
		}
		b.WriteString("\n")
	}

	if form.BaseURL != "" {
		b.WriteString("**Base URL:** " + form.BaseURL + "\n")
	}

	b.WriteString("\n")

	writeAuthSection(&b, form.Auth)
	writeGenerateSection(&b, opts)
	writeUsageSection(&b, opts.SchemaPath)

	leaves := collectLeaves(form.Commands, nil)
	examples := pickExamples(leaves, opts.SchemaPath, maxExampleCommands)
	if len(examples) > 0 {
		b.WriteString("## Example commands\n\n")
		for _, ex := range examples {
			if ex.cmd.Description != "" {
				b.WriteString(ex.cmd.Description + "\n\n")
			}
			b.WriteString("```bash\n" + ex.line + "\n```\n\n")
		}
	}

	b.WriteString("## Commands\n\n")
	if len(leaves) <= maxFullCommandDocs {
		writeCommandTree(&b, form.Commands, nil, opts.SchemaPath, true)
	} else {
		b.WriteString(fmt.Sprintf(
			"This CLI defines %d commands. Use `--help` to explore the full tree.\n\n",
			len(leaves),
		))
		writeCommandTree(&b, form.Commands, nil, opts.SchemaPath, false)
	}

	return b.String()
}

func formatLabel(format string) string {
	switch format {
	case "openapi":
		return "OpenAPI 3.x / Swagger"
	case "proto":
		return "gRPC / Protocol Buffers"
	case "graphql":
		return "GraphQL SDL"
	case "asyncapi":
		return "AsyncAPI 2.x / 3.x"
	default:
		return format
	}
}

func writeAuthSection(b *strings.Builder, auth AuthConfig) {
	b.WriteString("## Authentication\n\n")
	b.WriteString("**Type:** `" + auth.Type + "`\n")

	if auth.Comment != "" {
		b.WriteString("\n" + auth.Comment + "\n")
	}

	if len(auth.Config) > 0 {
		b.WriteString("\n| Config key | Source |\n")
		b.WriteString("|---|---|\n")
		keys := sortedStringKeys(auth.Config)
		for _, key := range keys {
			b.WriteString("| `" + key + "` | " + envOrValueSource(auth.Config[key]) + " |\n")
		}
	}
	b.WriteString("\n")
}

func envOrValueSource(v EnvOrValue) string {
	if v.Env != "" {
		return "env `" + v.Env + "`"
	}
	if v.File != "" {
		return "file `" + v.File + "`"
	}
	if v.Value != "" {
		return "literal"
	}
	return "—"
}

func writeGenerateSection(b *strings.Builder, opts DocOptions) {
	b.WriteString("## Generate\n\n")
	b.WriteString("```bash\n")
	b.WriteString("go run ./cmd/builder generate")
	if opts.SpecPath != "" {
		b.WriteString(" --spec " + opts.SpecPath)
	}
	if opts.SchemaPath != "" {
		b.WriteString(" --out " + opts.SchemaPath)
	}
	if opts.FormatForced && opts.Format != "" {
		b.WriteString(" --format " + opts.Format)
	}
	b.WriteString("\n```\n\n")
}

func writeUsageSection(b *strings.Builder, schemaPath string) {
	b.WriteString("## Usage\n\n")
	b.WriteString("```bash\n")
	b.WriteString("go run ./cmd/runner --form " + schemaPath + " --help\n")
	b.WriteString("```\n\n")
}

type leafEntry struct {
	path []string
	cmd  Command
}

func collectLeaves(commands map[string]Command, prefix []string) []leafEntry {
	var leaves []leafEntry
	for _, name := range sortedStringKeys(commands) {
		cmd := commands[name]
		path := append(prefix, name)
		if cmd.Action != nil {
			leaves = append(leaves, leafEntry{path: path, cmd: cmd})
		}
		if len(cmd.Commands) > 0 {
			leaves = append(leaves, collectLeaves(cmd.Commands, path)...)
		}
	}
	return leaves
}

func pickExamples(leaves []leafEntry, schemaPath string, max int) []struct {
	cmd  Command
	line string
} {
	priority := func(verb string) int {
		switch verb {
		case "list":
			return 0
		case "get":
			return 1
		case "create":
			return 2
		case "subscribe":
			return 3
		case "publish":
			return 4
		default:
			return 5
		}
	}

	sorted := append([]leafEntry{}, leaves...)
	sort.Slice(sorted, func(i, j int) bool {
		vi := priority(sorted[i].path[len(sorted[i].path)-1])
		vj := priority(sorted[j].path[len(sorted[j].path)-1])
		if vi != vj {
			return vi < vj
		}
		return strings.Join(sorted[i].path, " ") < strings.Join(sorted[j].path, " ")
	})

	var out []struct {
		cmd  Command
		line string
	}
	seen := make(map[string]bool)
	for _, leaf := range sorted {
		if len(out) >= max {
			break
		}
		key := strings.Join(leaf.path, "/")
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, struct {
			cmd  Command
			line string
		}{
			cmd:  leaf.cmd,
			line: buildExampleLine(schemaPath, leaf.path, leaf.cmd),
		})
	}
	return out
}

func buildExampleLine(schemaPath string, path []string, cmd Command) string {
	parts := []string{"go run ./cmd/runner", "--form", schemaPath}
	parts = append(parts, path...)
	for _, a := range cmd.Args {
		if a.Required {
			parts = append(parts, "<"+a.Name+">")
		}
	}
	for _, p := range cmd.Parameters {
		flag := strings.TrimPrefix(p.Flag, "--")
		if flag == "limit" {
			parts = append(parts, "--limit", "5")
			break
		}
	}
	return strings.Join(parts, " ")
}

func writeCommandTree(b *strings.Builder, commands map[string]Command, prefix []string, schemaPath string, detail bool) {
	for _, name := range sortedStringKeys(commands) {
		cmd := commands[name]
		path := append(prefix, name)
		heading := strings.Repeat("#", min(len(path)+2, 6))
		b.WriteString(heading + " " + name + "\n\n")
		if cmd.Description != "" {
			b.WriteString(cmd.Description + "\n\n")
		}

		if cmd.Action != nil && detail {
			b.WriteString("```bash\n" + buildExampleLine(schemaPath, path, cmd) + "\n```\n\n")
		}

		if len(cmd.Commands) > 0 {
			if !detail {
				count := countLeavesUnder(cmd.Commands)
				if count > 0 {
					b.WriteString(fmt.Sprintf("%d subcommand(s).\n\n", count))
				}
			}
			writeCommandTree(b, cmd.Commands, path, schemaPath, detail)
		}
	}
}

func countLeavesUnder(commands map[string]Command) int {
	return len(collectLeaves(commands, nil))
}

func sortedStringKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
