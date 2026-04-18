// Package formats detects the API spec format from a file path and content.
package formats

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const (
	OpenAPI  = "openapi"
	Proto    = "proto"
	GraphQL  = "graphql"
	AsyncAPI = "asyncapi"
	Unknown  = "unknown"
)

// Detect returns the format constant for the given file.
// It first checks the file extension, then peeks at the content for YAML/JSON files.
func Detect(path string) string {
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".proto":
		return Proto
	case ".graphql", ".gql":
		return GraphQL
	}

	// For JSON/YAML, peek at top-level keys.
	if ext == ".json" || ext == ".yaml" || ext == ".yml" {
		return detectFromContent(path)
	}

	return Unknown
}

func detectFromContent(path string) string {
	f, err := os.Open(path)
	if err != nil {
		return Unknown
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := 0
	for scanner.Scan() && lines < 20 {
		line := strings.TrimSpace(scanner.Text())
		lines++

		if strings.HasPrefix(line, `"openapi"`) || strings.HasPrefix(line, "openapi:") ||
			strings.HasPrefix(line, `"swagger"`) || strings.HasPrefix(line, "swagger:") {
			return OpenAPI
		}
		if strings.HasPrefix(line, "asyncapi:") || strings.HasPrefix(line, `"asyncapi"`) {
			return AsyncAPI
		}
		// GraphQL SDL in YAML/JSON is unusual but handle it.
		if strings.HasPrefix(line, "type Query") || strings.HasPrefix(line, "type Mutation") {
			return GraphQL
		}
	}
	return Unknown
}
