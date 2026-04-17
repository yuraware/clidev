package oas

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Parse reads an OpenAPI 3.x JSON or YAML file and returns a flat list of operations.
func Parse(path string) (*Spec, []Op, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("reading spec: %w", err)
	}

	// The ACS spec is JSON; add YAML support if needed later.
	var spec Spec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, nil, fmt.Errorf("parsing spec JSON: %w", err)
	}

	ops := flattenOps(&spec)
	return &spec, ops, nil
}

func flattenOps(spec *Spec) []Op {
	var ops []Op
	for path, item := range spec.Paths {
		for _, pair := range []struct {
			method string
			op     *Operation
		}{
			{"GET", item.Get},
			{"POST", item.Post},
			{"PATCH", item.Patch},
			{"DELETE", item.Delete},
			{"PUT", item.Put},
		} {
			if pair.op == nil {
				continue
			}
			o := pair.op
			ops = append(ops, Op{
				OperationID: o.OperationID,
				Method:      pair.method,
				Path:        path,
				Tags:        o.Tags,
				Parameters:  resolveParams(o.Parameters),
				HasBody:     o.RequestBody != nil,
			})
		}
	}
	return ops
}

// resolveParams filters out any $ref parameters (we don't need to inline them for CLI generation).
func resolveParams(params []Parameter) []Parameter {
	out := make([]Parameter, 0, len(params))
	for _, p := range params {
		if p.Name != "" {
			out = append(out, p)
		}
	}
	return out
}

// TagsFromPath returns the primary tag for an operation, defaulting to the path's first segment.
func TagsFromPath(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) >= 2 {
		return parts[1] // e.g. /v1/apps → apps
	}
	return "misc"
}
