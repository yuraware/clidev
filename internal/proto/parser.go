// Package proto parses Protocol Buffer .proto files and generates a CLIForm.
// It extracts services, RPCs, and google.api.http annotations for HTTP transcoding.
package proto

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Spec is the parsed representation of a .proto file.
type Spec struct {
	PackageName string
	Services    []Service
	Messages    map[string]Message
	DefaultHost string
	OAuthScopes []string
}

// Service is a proto service block.
type Service struct {
	Name    string
	RPCs    []RPC
	Comment string
}

// RPC is a single remote procedure call.
type RPC struct {
	Name       string
	InputType  string
	OutputType string
	HTTPMethod string // GET POST PUT PATCH DELETE
	HTTPPath   string
	Body       string // "*" or field name
	Comment    string
	Streaming  bool
}

// Message is a proto message type.
type Message struct {
	Name   string
	Fields []Field
}

// Field is a message field.
type Field struct {
	Name     string
	Type     string
	Number   int
	Repeated bool
	Required bool // field_behavior = REQUIRED
}

var (
	reService     = regexp.MustCompile(`^service\s+(\w+)`)
	reRPC         = regexp.MustCompile(`^\s*rpc\s+(\w+)\s*\(([^)]+)\)\s+returns\s+\(([^)]+)\)`)
	reHTTPMethod  = regexp.MustCompile(`^\s*(get|post|put|patch|delete):\s*"([^"]+)"`)
	reHTTPBody    = regexp.MustCompile(`^\s*body:\s*"([^"]*)"`)
	reDefaultHost = regexp.MustCompile(`option\s+\(google\.api\.default_host\)\s*=\s*"([^"]+)"`)
	reOAuthScopes = regexp.MustCompile(`option\s+\(google\.api\.oauth_scopes\)\s*=\s*"([^"]+)"`)
	reMessage     = regexp.MustCompile(`^message\s+(\w+)`)
	reField       = regexp.MustCompile(`^\s*(repeated\s+)?(\w[\w.]*)\s+(\w+)\s*=\s*(\d+)`)
)

// Parse reads a .proto file and returns a Spec.
func Parse(path string) (*Spec, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening proto file: %w", err)
	}
	defer f.Close()

	spec := &Spec{Messages: make(map[string]Message)}
	scanner := bufio.NewScanner(f)

	var (
		inService    bool
		inRPC        bool
		inHTTP       bool
		inMessage    bool
		currentSvc   Service
		currentRPC   RPC
		currentMsg   Message
		commentBuf   strings.Builder
		depth        int
	)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// Collect block comment lines.
		if strings.HasPrefix(trimmed, "//") {
			commentBuf.WriteString(strings.TrimPrefix(strings.TrimPrefix(trimmed, "//"), " "))
			commentBuf.WriteString(" ")
			continue
		}

		// Package.
		if strings.HasPrefix(trimmed, "package ") {
			spec.PackageName = strings.TrimSuffix(strings.TrimPrefix(trimmed, "package "), ";")
		}

		// Default host (first occurrence, inside first service).
		if m := reDefaultHost.FindStringSubmatch(trimmed); m != nil {
			if spec.DefaultHost == "" {
				spec.DefaultHost = m[1]
			}
		}

		// OAuth scopes.
		if m := reOAuthScopes.FindStringSubmatch(trimmed); m != nil {
			for _, s := range strings.Split(m[1], ",") {
				s = strings.TrimSpace(s)
				if s != "" {
					spec.OAuthScopes = append(spec.OAuthScopes, s)
				}
			}
		}

		// Service start.
		if !inService && !inMessage {
			if m := reService.FindStringSubmatch(trimmed); m != nil {
				inService = true
				depth = 0
				currentSvc = Service{Name: m[1], Comment: strings.TrimSpace(commentBuf.String())}
				commentBuf.Reset()
				if strings.Contains(trimmed, "{") {
					depth++
				}
				continue
			}
		}

		// Message start.
		if !inService && !inMessage {
			if m := reMessage.FindStringSubmatch(trimmed); m != nil {
				inMessage = true
				depth = 0
				currentMsg = Message{Name: m[1]}
				if strings.Contains(trimmed, "{") {
					depth++
				}
				commentBuf.Reset()
				continue
			}
		}

		commentBuf.Reset()

		// Track brace depth.
		depth += strings.Count(trimmed, "{") - strings.Count(trimmed, "}")

		if inMessage {
			if depth <= 0 {
				spec.Messages[currentMsg.Name] = currentMsg
				inMessage = false
				continue
			}
			if m := reField.FindStringSubmatch(line); m != nil {
				field := Field{
					Name:     m[3],
					Type:     m[2],
					Repeated: m[1] != "",
				}
				currentMsg.Fields = append(currentMsg.Fields, field)
			}
			continue
		}

		if inService {
			if depth <= 0 {
				spec.Services = append(spec.Services, currentSvc)
				inService = false
				inRPC = false
				continue
			}

			// RPC declaration.
			if !inRPC {
				if m := reRPC.FindStringSubmatch(line); m != nil {
					inRPC = true
					inHTTP = false
					inputType := strings.TrimPrefix(strings.TrimSpace(m[2]), "stream ")
					outputType := strings.TrimPrefix(strings.TrimSpace(m[3]), "stream ")
					currentRPC = RPC{
						Name:      m[1],
						InputType: strings.TrimSpace(inputType),
						OutputType: strings.TrimSpace(outputType),
						Streaming: strings.Contains(m[2], "stream") || strings.Contains(m[3], "stream"),
					}
				}
				continue
			}

			// Inside RPC — look for HTTP annotation.
			if strings.Contains(trimmed, "google.api.http") {
				inHTTP = true
				continue
			}

			if inHTTP {
				if m := reHTTPMethod.FindStringSubmatch(line); m != nil {
					currentRPC.HTTPMethod = strings.ToUpper(m[1])
					currentRPC.HTTPPath = cleanHTTPPath(m[2])
				}
				if m := reHTTPBody.FindStringSubmatch(line); m != nil {
					currentRPC.Body = m[1]
				}
			}

			// End of RPC block.
			if inRPC && strings.Contains(trimmed, "}") && depth == 1 {
				if !currentRPC.Streaming {
					currentSvc.RPCs = append(currentSvc.RPCs, currentRPC)
				}
				inRPC = false
				inHTTP = false
			}
		}
	}

	return spec, scanner.Err()
}

// cleanHTTPPath normalises proto HTTP path patterns to simple {param} style.
// e.g. "/v1/{name=projects/*/topics/*}" → "/v1/{name}"
func cleanHTTPPath(path string) string {
	re := regexp.MustCompile(`\{(\w+)=[^}]+\}`)
	return re.ReplaceAllString(path, "{$1}")
}
