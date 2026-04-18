package cliSchema

// CLIForm is the root of a cli-schema.yaml / cli-schema.json document.
type CLIForm struct {
	Name        string             `yaml:"name"        json:"name"`
	Version     string             `yaml:"version"     json:"version"`
	Description string             `yaml:"description" json:"description"`
	BaseURL     string             `yaml:"base_url"    json:"base_url"`
	Auth        AuthConfig         `yaml:"auth"        json:"auth"`
	Output      OutputConfig       `yaml:"output"      json:"output"`
	Commands    map[string]Command `yaml:"commands"    json:"commands"`
}

// AuthConfig describes how to authenticate requests.
//
// Types:
//
//	none        — no authentication
//	bearer      — Authorization: Bearer <token>  (generic / GitHub PAT)
//	api_key     — custom header or query param key  (config: key, header, in)
//	basic       — Authorization: Basic base64(user:pass)  (config: username, password)
//	oauth2      — client-credentials token exchange  (config: client_id, client_secret, token_url, scopes)
//	apple_jwt   — Apple-signed ES256 JWT  (config: key_id, issuer_id, private_key, audience)
type AuthConfig struct {
	Type    string                `yaml:"type"              json:"type"`
	Config  map[string]EnvOrValue `yaml:"config,omitempty"  json:"config,omitempty"`
	Comment string                `yaml:"comment,omitempty" json:"comment,omitempty"`
}

// EnvOrValue resolves a config value from an env var, literal value, or file path.
type EnvOrValue struct {
	Env   string `yaml:"env,omitempty"   json:"env,omitempty"`
	Value string `yaml:"value,omitempty" json:"value,omitempty"`
	File  string `yaml:"file,omitempty"  json:"file,omitempty"`
}

// OutputConfig controls default output behaviour.
type OutputConfig struct {
	// DefaultFormat is one of: table | json | yaml
	DefaultFormat string `yaml:"default_format" json:"default_format"`
}

// Command is either a group (has sub-Commands) or a leaf (has Action).
type Command struct {
	Description string             `yaml:"description"          json:"description"`
	Action      *Action            `yaml:"action,omitempty"     json:"action,omitempty"`
	Args        []Arg              `yaml:"args,omitempty"       json:"args,omitempty"`
	Parameters  []Parameter        `yaml:"parameters,omitempty" json:"parameters,omitempty"`
	Body        *BodyConfig        `yaml:"body,omitempty"       json:"body,omitempty"`
	Commands    map[string]Command `yaml:"commands,omitempty"   json:"commands,omitempty"`
}

// Action describes the HTTP operation a leaf command performs.
//
// Type is one of: http (default) | graphql
//   - http:    standard REST call; Method + Path are used.
//   - graphql: POSTs a GraphQL document to Path; GraphQLQuery holds the operation.
type Action struct {
	Type         string `yaml:"type,omitempty"          json:"type,omitempty"`
	Method       string `yaml:"method"                  json:"method"`
	Path         string `yaml:"path"                    json:"path"`
	OperationID  string `yaml:"operation_id,omitempty"  json:"operation_id,omitempty"`
	GraphQLQuery string `yaml:"graphql_query,omitempty" json:"graphql_query,omitempty"`
}

// Parameter maps a CLI flag to an HTTP query parameter.
type Parameter struct {
	Flag    string   `yaml:"flag"                  json:"flag"`
	Query   string   `yaml:"query"                 json:"query"`
	// Type is one of: string | string_array | integer | boolean | enum | enum_array
	Type    string   `yaml:"type"                  json:"type"`
	Values  []string `yaml:"values,omitempty"      json:"values,omitempty"`
	Default any      `yaml:"default,omitempty"     json:"default,omitempty"`
	Max     int      `yaml:"max,omitempty"         json:"max,omitempty"`
	Desc    string   `yaml:"description,omitempty" json:"description,omitempty"`
}

// Arg is a positional CLI argument that maps to a path parameter or GraphQL variable.
type Arg struct {
	Name      string `yaml:"name"                   json:"name"`
	PathParam string `yaml:"path_param,omitempty"   json:"path_param,omitempty"`
	Required  bool   `yaml:"required"               json:"required"`
	Desc      string `yaml:"description,omitempty"  json:"description,omitempty"`
}

// BodyConfig describes how to build a request body from CLI flags.
type BodyConfig struct {
	// Format is one of: json_api | raw_json | graphql_vars
	Format        string         `yaml:"format"                  json:"format"`
	ResourceType  string         `yaml:"resource_type,omitempty" json:"resource_type,omitempty"`
	Attributes    []BodyField    `yaml:"attributes,omitempty"    json:"attributes,omitempty"`
	Relationships []BodyRelField `yaml:"relationships,omitempty" json:"relationships,omitempty"`
}

// BodyField maps a CLI flag to a request body field.
type BodyField struct {
	Flag     string `yaml:"flag"               json:"flag"`
	Field    string `yaml:"field"              json:"field"`
	Type     string `yaml:"type"               json:"type"`
	Required bool   `yaml:"required,omitempty" json:"required,omitempty"`
	Default  any    `yaml:"default,omitempty"  json:"default,omitempty"`
	Desc     string `yaml:"description,omitempty" json:"description,omitempty"`
}

// BodyRelField maps a CLI flag to a JSON:API relationship.
type BodyRelField struct {
	Flag         string `yaml:"flag"               json:"flag"`
	Relationship string `yaml:"relationship"       json:"relationship"`
	ResourceType string `yaml:"resource_type"      json:"resource_type"`
	Required     bool   `yaml:"required,omitempty" json:"required,omitempty"`
}
