package cliSchema

import (
	"strings"
	"testing"
)

func TestGenerateMarkdown_containsCoreSections(t *testing.T) {
	form := &CLIForm{
		Name:        "streetlights",
		Version:     "1.0",
		Description: "Streetlights Kafka API",
		BaseURL:     "test.mykafkaserver.com:9092",
		Auth: AuthConfig{
			Type: "basic",
			Config: map[string]EnvOrValue{
				"username": {Env: "KAFKA_USERNAME"},
				"password": {Env: "KAFKA_PASSWORD"},
			},
		},
		Commands: map[string]Command{
			"lightturnon": {
				Description: "Turn on commands",
				Commands: map[string]Command{
					"publish": {
						Description: "Turn on a streetlight",
						Action: &Action{
							Method: "POST",
							Path:   "/publish",
						},
					},
				},
			},
		},
	}

	md := GenerateMarkdown(form, DocOptions{
		SpecPath:   "sample-api/streetlights/asyncapi.yml",
		SchemaPath: "sample-api/streetlights/cli-schema.yaml",
		Format:     "asyncapi",
		Stats:      "4 channels · 4 operations",
	})

	checks := []string{
		"# Streetlights Kafka API",
		"**Spec format:** AsyncAPI 2.x / 3.x · 4 channels · 4 operations",
		"## Authentication",
		"`basic`",
		"KAFKA_USERNAME",
		"## Generate",
		"--spec sample-api/streetlights/asyncapi.yml",
		"--out sample-api/streetlights/cli-schema.yaml",
		"## Example commands",
		"lightturnon publish",
		"## Commands",
	}
	for _, want := range checks {
		if !strings.Contains(md, want) {
			t.Errorf("markdown missing %q\n%s", want, md)
		}
	}
}

func TestDocPathFor(t *testing.T) {
	got := DocPathFor("sample-api/acs/cli-schema.yaml")
	want := "sample-api/acs/cli-schema.md"
	if got != want {
		t.Errorf("DocPathFor() = %q, want %q", got, want)
	}
}
