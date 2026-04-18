// Package asyncapi parses AsyncAPI 2.x/3.x spec files and generates a CLIForm.
package asyncapi

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/yurikobets/cli-builder/internal/cliSchema"
)

// Spec is the parsed representation of an AsyncAPI document.
type Spec struct {
	AsyncAPI   string             `yaml:"asyncapi"`
	Info       Info               `yaml:"info"`
	Servers    map[string]Server  `yaml:"servers"`
	Channels   map[string]Channel `yaml:"channels"`
	Operations map[string]Op      `yaml:"operations"`
	Components Components         `yaml:"components"`
}

type Info struct {
	Title       string `yaml:"title"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
}

type Server struct {
	Host        string            `yaml:"host"`
	Protocol    string            `yaml:"protocol"`
	Description string            `yaml:"description"`
	Security    []map[string]any  `yaml:"security"`
}

type Channel struct {
	Address     string              `yaml:"address"`
	Description string              `yaml:"description"`
	Messages    map[string]Ref      `yaml:"messages"`
	Parameters  map[string]Param    `yaml:"parameters"`
}

type Op struct {
	Action  string         `yaml:"action"` // send | receive (3.x) or publish | subscribe (2.x)
	Channel Ref            `yaml:"channel"`
	Message Ref            `yaml:"message"`
	Summary string         `yaml:"summary"`
}

type Ref struct {
	Ref string `yaml:"$ref"`
}

type Param struct {
	Description string     `yaml:"description"`
	Schema      SchemaNode `yaml:"schema"`
}

type Components struct {
	Messages        map[string]Message        `yaml:"messages"`
	Schemas         map[string]SchemaNode     `yaml:"schemas"`
	SecuritySchemes map[string]SecurityScheme `yaml:"securitySchemes"`
}

type Message struct {
	Name        string     `yaml:"name"`
	Title       string     `yaml:"title"`
	Summary     string     `yaml:"summary"`
	Description string     `yaml:"description"`
	Payload     SchemaNode `yaml:"payload"`
}

type SchemaNode struct {
	Type       string                `yaml:"type"`
	Properties map[string]SchemaNode `yaml:"properties"`
	Required   []string              `yaml:"required"`
	Ref        string                `yaml:"$ref"`
}

type SecurityScheme struct {
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
}

// Parse reads an AsyncAPI YAML/JSON file and returns a Spec.
func Parse(path string) (*Spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading asyncapi spec: %w", err)
	}

	var spec Spec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("parsing asyncapi spec: %w", err)
	}

	if spec.AsyncAPI == "" {
		return nil, fmt.Errorf("not an AsyncAPI document (missing 'asyncapi' key)")
	}

	return &spec, nil
}

// Generate converts a parsed AsyncAPI Spec into a CLIForm.
func Generate(spec *Spec) *cliSchema.CLIForm {
	baseURL := serverURL(spec)

	form := &cliSchema.CLIForm{
		Name:        slugify(spec.Info.Title),
		Version:     spec.Info.Version,
		Description: spec.Info.Title,
		BaseURL:     baseURL,
		Auth:        inferAuth(spec),
		Output:      cliSchema.OutputConfig{DefaultFormat: "json"},
		Commands:    make(map[string]cliSchema.Command),
	}

	// AsyncAPI 3.x: operations reference channels.
	if len(spec.Operations) > 0 {
		for opName, op := range spec.Operations {
			channelRef := deref(op.Channel.Ref)
			channel, ok := spec.Channels[channelRef]
			if !ok {
				continue
			}
			verb := actionToVerb(op.Action)
			grpName := slugify(channelRef)

			grp := form.Commands[grpName]
			if grp.Commands == nil {
				grp = cliSchema.Command{
					Description: channel.Description,
					Commands:    make(map[string]cliSchema.Command),
				}
			}

			cmd := buildOpCommand(opName, op, channel, spec)
			grp.Commands[verb] = cmd
			form.Commands[grpName] = grp
		}
		return form
	}

	// AsyncAPI 2.x fallback: channels have publish/subscribe inline.
	for chanName, channel := range spec.Channels {
		grpName := slugify(chanName)
		grp := cliSchema.Command{
			Description: channel.Description,
			Commands:    make(map[string]cliSchema.Command),
		}

		// Channel address parameters → positional args.
		var args []cliSchema.Arg
		for paramName, p := range channel.Parameters {
			args = append(args, cliSchema.Arg{
				Name:     paramName,
				Required: true,
				Desc:     p.Description,
			})
		}

		msg := resolveFirstMessage(channel, spec)
		body := buildBodyFromSchema(msg)

		grp.Commands["publish"] = cliSchema.Command{
			Description: fmt.Sprintf("Publish a message to %s", channel.Address),
			Action: &cliSchema.Action{
				Method:      "POST",
				Path:        channel.Address,
				OperationID: chanName + "_publish",
			},
			Args: args,
			Body: body,
		}
		grp.Commands["subscribe"] = cliSchema.Command{
			Description: fmt.Sprintf("Subscribe to messages from %s", channel.Address),
			Action: &cliSchema.Action{
				Method:      "GET",
				Path:        channel.Address,
				OperationID: chanName + "_subscribe",
			},
			Args: args,
		}
		form.Commands[grpName] = grp
	}

	return form
}

func buildOpCommand(opName string, op Op, channel Channel, spec *Spec) cliSchema.Command {
	var args []cliSchema.Arg
	for paramName, p := range channel.Parameters {
		args = append(args, cliSchema.Arg{
			Name:     paramName,
			Required: true,
			Desc:     p.Description,
		})
	}

	desc := op.Summary
	if desc == "" {
		desc = fmt.Sprintf("%s on %s", op.Action, channel.Address)
	}

	msg := resolveFirstMessage(channel, spec)
	var body *cliSchema.BodyConfig
	if op.Action == "send" || op.Action == "publish" {
		body = buildBodyFromSchema(msg)
	}

	method := "POST"
	if op.Action == "receive" || op.Action == "subscribe" {
		method = "GET"
	}

	return cliSchema.Command{
		Description: desc,
		Action: &cliSchema.Action{
			Method:      method,
			Path:        channel.Address,
			OperationID: opName,
		},
		Args: args,
		Body: body,
	}
}

func resolveFirstMessage(channel Channel, spec *Spec) *Message {
	for _, ref := range channel.Messages {
		name := deref(ref.Ref)
		if msg, ok := spec.Components.Messages[name]; ok {
			return &msg
		}
	}
	return nil
}

func buildBodyFromSchema(msg *Message) *cliSchema.BodyConfig {
	if msg == nil {
		return nil
	}
	body := &cliSchema.BodyConfig{Format: "raw_json"}
	for fieldName, prop := range msg.Payload.Properties {
		body.Attributes = append(body.Attributes, cliSchema.BodyField{
			Flag:  "--" + fieldName,
			Field: fieldName,
			Type:  jsonTypeToParam(prop.Type),
			Desc:  "",
		})
	}
	if len(body.Attributes) == 0 {
		return nil
	}
	return body
}

func inferAuth(spec *Spec) cliSchema.AuthConfig {
	for _, scheme := range spec.Components.SecuritySchemes {
		switch strings.ToLower(scheme.Type) {
		case "scramsha256", "scramsha512":
			return cliSchema.AuthConfig{
				Type: "basic",
				Config: map[string]cliSchema.EnvOrValue{
					"username": {Env: "KAFKA_USERNAME"},
					"password": {Env: "KAFKA_PASSWORD"},
				},
				Comment: "SASL/SCRAM — set KAFKA_USERNAME and KAFKA_PASSWORD",
			}
		case "x509", "mtls":
			return cliSchema.AuthConfig{
				Type:    "none",
				Comment: "mTLS — configure client certificates at the transport layer",
			}
		case "httpbearer", "http":
			return cliSchema.AuthConfig{
				Type: "bearer",
				Config: map[string]cliSchema.EnvOrValue{
					"token": {Env: "API_TOKEN"},
				},
			}
		case "apikey":
			return cliSchema.AuthConfig{
				Type: "api_key",
				Config: map[string]cliSchema.EnvOrValue{
					"key":    {Env: "API_KEY"},
					"header": {Value: "X-API-Key"},
					"in":     {Value: "header"},
				},
			}
		}
	}
	return cliSchema.AuthConfig{Type: "none"}
}

func serverURL(spec *Spec) string {
	for _, s := range spec.Servers {
		if s.Host != "" {
			protocol := s.Protocol
			if protocol == "" {
				protocol = "https"
			}
			if !strings.Contains(s.Host, "://") {
				return protocol + "://" + s.Host
			}
			return s.Host
		}
	}
	return ""
}

func actionToVerb(action string) string {
	switch action {
	case "send", "publish":
		return "publish"
	case "receive", "subscribe":
		return "subscribe"
	}
	return action
}

// deref strips a JSON Pointer like "#/channels/foo" to "foo".
func deref(ref string) string {
	parts := strings.Split(ref, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ref
}

func jsonTypeToParam(t string) string {
	switch t {
	case "integer", "number":
		return "integer"
	case "boolean":
		return "boolean"
	case "array":
		return "string_array"
	default:
		return "string"
	}
}

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.NewReplacer(" ", "-", "_", "-", "/", "-", ".", "-").Replace(s)
	return strings.Trim(s, "-")
}
