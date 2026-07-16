package runtime

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yuraware/clidev/internal/cliSchema"
	"github.com/yuraware/clidev/internal/output"
	"github.com/yuraware/clidev/internal/runtime/auth"
)

// Build constructs a root *cobra.Command from a CLIForm.
func Build(form *cliSchema.CLIForm) (*cobra.Command, error) {
	authProvider, err := auth.New(form.Auth)
	if err != nil {
		return nil, fmt.Errorf("initialising auth: %w", err)
	}

	exec := NewExecutor(form.BaseURL, authProvider)
	outFormat := form.Output.DefaultFormat
	if outFormat == "" {
		outFormat = "table"
	}

	root := &cobra.Command{
		Use:   form.Name,
		Short: form.Description,
		Long:  fmt.Sprintf("%s (v%s)\n%s", form.Name, form.Version, form.Description),
	}

	// Global --output flag.
	root.PersistentFlags().StringP("output", "o", outFormat, "Output format: table | json")

	// Sort resource names for deterministic help output.
	resources := sortedKeys(form.Commands)
	for _, resName := range resources {
		resCmd := form.Commands[resName]
		root.AddCommand(buildGroup(resName, resCmd, exec, root))
	}

	return root, nil
}

func buildGroup(name string, cmd cliSchema.Command, exec *Executor, root *cobra.Command) *cobra.Command {
	grp := &cobra.Command{
		Use:   name,
		Short: cmd.Description,
	}

	// If this node itself has an action it's a leaf too.
	if cmd.Action != nil {
		wireAction(grp, cmd, exec, root)
	}

	subNames := sortedKeys(cmd.Commands)
	for _, subName := range subNames {
		sub := cmd.Commands[subName]
		grp.AddCommand(buildGroup(subName, sub, exec, root))
	}

	return grp
}

// wireAction attaches RunE + flags to a Cobra command for a leaf action.
func wireAction(cobraCmd *cobra.Command, cmd cliSchema.Command, exec *Executor, root *cobra.Command) {
	// Register flags for all parameters.
	for _, p := range cmd.Parameters {
		flagName := strings.TrimPrefix(p.Flag, "--")
		desc := p.Desc
		if desc == "" {
			desc = fmt.Sprintf("query param: %s", p.Query)
		}
		if p.Max > 0 {
			desc = fmt.Sprintf("%s (max %d)", desc, p.Max)
		}
		if len(p.Values) > 0 {
			desc = fmt.Sprintf("%s [%s]", desc, strings.Join(p.Values, "|"))
		}

		switch p.Type {
		case "integer":
			def := 0
			if p.Default != nil {
				switch v := p.Default.(type) {
				case int:
					def = v
				case float64:
					def = int(v)
				case string:
					def, _ = strconv.Atoi(v)
				}
			}
			cobraCmd.Flags().Int(flagName, def, desc)
		case "boolean":
			cobraCmd.Flags().Bool(flagName, false, desc)
		default:
			// string, string_array, enum, enum_array — all stored as strings.
			def := ""
			if p.Default != nil {
				def = fmt.Sprintf("%v", p.Default)
			}
			cobraCmd.Flags().String(flagName, def, desc)
		}
	}

	// Register flags for body attributes.
	if cmd.Body != nil {
		for _, f := range cmd.Body.Attributes {
			flagName := strings.TrimPrefix(f.Flag, "--")
			desc := f.Desc
			if f.Required {
				desc = "(required) " + desc
			}
			def := ""
			if f.Default != nil {
				def = fmt.Sprintf("%v", f.Default)
			}
			cobraCmd.Flags().String(flagName, def, desc)
		}
		for _, r := range cmd.Body.Relationships {
			flagName := strings.TrimPrefix(r.Flag, "--")
			desc := fmt.Sprintf("ID of related %s", r.ResourceType)
			if r.Required {
				desc = "(required) " + desc
			}
			cobraCmd.Flags().String(flagName, "", desc)
		}
	}

	cobraCmd.RunE = func(cobraCmd *cobra.Command, args []string) error {
		// Validate arg count.
		if len(args) < requiredArgCount(cmd) {
			return fmt.Errorf("expected %d argument(s): %s", requiredArgCount(cmd), argUsage(cmd))
		}

		respBody, err := exec.Execute(cobraCmd, args, cmd)
		if err != nil {
			return err
		}

		// Resolve output format (local flag overrides root default).
		outFmt, _ := root.PersistentFlags().GetString("output")
		if f, err := cobraCmd.Flags().GetString("output"); err == nil && f != "" {
			outFmt = f
		}

		return output.Format(os.Stdout, respBody, outFmt)
	}

	// Build Use string with args.
	if len(cmd.Args) > 0 {
		cobraCmd.Use = cobraCmd.Use + " " + argUsage(cmd)
	}
}

func requiredArgCount(cmd cliSchema.Command) int {
	n := 0
	for _, a := range cmd.Args {
		if a.Required {
			n++
		}
	}
	return n
}

func argUsage(cmd cliSchema.Command) string {
	parts := make([]string, len(cmd.Args))
	for i, a := range cmd.Args {
		if a.Required {
			parts[i] = "<" + a.Name + ">"
		} else {
			parts[i] = "[" + a.Name + "]"
		}
	}
	return strings.Join(parts, " ")
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
