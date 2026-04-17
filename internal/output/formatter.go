package output

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Format prints the API response body in the requested format.
func Format(w io.Writer, data []byte, format string) error {
	switch format {
	case "json":
		return printJSON(w, data)
	default: // table
		return printTable(w, data)
	}
}

func printJSON(w io.Writer, data []byte) error {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		_, err = fmt.Fprintln(w, string(data))
		return err
	}
	pretty, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(pretty))
	return err
}

// printTable handles JSON:API single resource and collection responses.
func printTable(w io.Writer, data []byte) error {
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(data, &envelope); err != nil {
		_, err = fmt.Fprintln(w, string(data))
		return err
	}

	rawData, ok := envelope["data"]
	if !ok {
		return printJSON(w, data)
	}

	// Detect array vs single object.
	trimmed := strings.TrimSpace(string(rawData))
	if strings.HasPrefix(trimmed, "[") {
		return printCollection(w, rawData)
	}
	return printSingle(w, rawData)
}

func printCollection(w io.Writer, rawData json.RawMessage) error {
	var items []map[string]json.RawMessage
	if err := json.Unmarshal(rawData, &items); err != nil {
		_, err = fmt.Fprintln(w, string(rawData))
		return err
	}
	if len(items) == 0 {
		fmt.Fprintln(w, "(no results)")
		return nil
	}

	// Collect all attribute keys across items for consistent columns.
	colSet := map[string]bool{"id": true, "type": true}
	for _, item := range items {
		if attrRaw, ok := item["attributes"]; ok {
			var attrs map[string]json.RawMessage
			if err := json.Unmarshal(attrRaw, &attrs); err == nil {
				for k := range attrs {
					colSet[k] = true
				}
			}
		}
	}

	cols := sortedKeys(colSet)
	colWidths := make([]int, len(cols))
	for i, c := range cols {
		colWidths[i] = len(c)
	}

	// Build rows.
	rows := make([][]string, len(items))
	for i, item := range items {
		row := make([]string, len(cols))
		flat := flattenItem(item)
		for j, c := range cols {
			v := flat[c]
			row[j] = v
			if len(v) > colWidths[j] {
				colWidths[j] = len(v)
			}
		}
		rows[i] = row
	}

	printTableRows(w, cols, rows, colWidths)
	fmt.Fprintf(w, "\n%d items\n", len(items))
	return nil
}

func printSingle(w io.Writer, rawData json.RawMessage) error {
	var item map[string]json.RawMessage
	if err := json.Unmarshal(rawData, &item); err != nil {
		_, err = fmt.Fprintln(w, string(rawData))
		return err
	}
	flat := flattenItem(item)
	keys := sortedKeys(flat)
	maxKey := 0
	for _, k := range keys {
		if len(k) > maxKey {
			maxKey = len(k)
		}
	}
	for _, k := range keys {
		fmt.Fprintf(w, "%-*s  %s\n", maxKey, k, flat[k])
	}
	return nil
}

// flattenItem merges id, type, and attributes into a single string map.
func flattenItem(item map[string]json.RawMessage) map[string]string {
	out := map[string]string{}
	for _, topKey := range []string{"id", "type"} {
		if v, ok := item[topKey]; ok {
			out[topKey] = unquote(string(v))
		}
	}
	if attrRaw, ok := item["attributes"]; ok {
		var attrs map[string]json.RawMessage
		if err := json.Unmarshal(attrRaw, &attrs); err == nil {
			for k, v := range attrs {
				out[k] = unquote(string(v))
			}
		}
	}
	return out
}

func printTableRows(w io.Writer, headers []string, rows [][]string, widths []int) {
	sep := make([]string, len(headers))
	for i, h := range headers {
		if widths[i] < len(h) {
			widths[i] = len(h)
		}
		sep[i] = strings.Repeat("-", widths[i])
	}

	printRow(w, headers, widths)
	printRow(w, sep, widths)
	for _, row := range rows {
		printRow(w, row, widths)
	}
}

func printRow(w io.Writer, cols []string, widths []int) {
	parts := make([]string, len(cols))
	for i, c := range cols {
		parts[i] = fmt.Sprintf("%-*s", widths[i], c)
	}
	fmt.Fprintln(w, strings.Join(parts, "  "))
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func unquote(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}
