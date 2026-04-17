package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yurikobets/cli-builder/internal/cliSchema"
	"github.com/yurikobets/cli-builder/internal/runtime"
)

func main() {
	// When invoked as `runner --form cli-schema.yaml <command>...`,
	// strip --form before handing off to the generated Cobra tree.
	formPath, remaining := extractFormFlag(os.Args[1:])

	if formPath == "" {
		// Show usage when no --form is given.
		fmt.Fprintf(os.Stderr, "Usage: runner --form <cli-schema.yaml> [command...]\n")
		os.Exit(1)
	}

	form, err := cliSchema.Load(formPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading cli-schema: %v\n", err)
		os.Exit(1)
	}

	root, err := runtime.Build(form)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building CLI: %v\n", err)
		os.Exit(1)
	}

	// Override os.Args so Cobra parses the remaining args.
	os.Args = append([]string{form.Name}, remaining...)

	root.SetErr(os.Stderr)
	root.SetOut(os.Stdout)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

// extractFormFlag pulls --form <path> out of args, returning the path and the rest.
func extractFormFlag(args []string) (string, []string) {
	// Support both `--form=path` and `--form path`.
	var formPath string
	var rest []string

	skipNext := false
	for i, a := range args {
		if skipNext {
			skipNext = false
			continue
		}
		if a == "--form" && i+1 < len(args) {
			formPath = args[i+1]
			skipNext = true
			continue
		}
		if len(a) > 7 && a[:7] == "--form=" {
			formPath = a[7:]
			continue
		}
		rest = append(rest, a)
	}

	// Also check RUNNER_FORM env var as fallback.
	if formPath == "" {
		formPath = os.Getenv("RUNNER_FORM")
	}

	return formPath, rest
}

// help prints a standalone help message when no cli-schema is available.
func help() *cobra.Command {
	return &cobra.Command{
		Use:   "runner",
		Short: "Execute a CLI defined by a cli-schema file",
		Long: `runner loads a cli-schema.yaml (or .json) file and dynamically builds
a fully-featured CLI from it, then executes the requested command.

Usage:
  runner --form <path/to/cli-schema.yaml> [command] [flags]
  RUNNER_FORM=cli-schema.yaml runner [command] [flags]

Examples:
  runner --form sample-api/acs/cli-schema.yaml apps list --limit 5
  runner --form sample-api/acs/cli-schema.yaml apps get <id>
  runner --form sample-api/acs/cli-schema.yaml builds list --filter-app <app-id>
`,
	}
}
