package oas

// Spec is the top-level OpenAPI 3.x document (fields we care about).
type Spec struct {
	Info    Info                `json:"info"`
	Servers []Server            `json:"servers"`
	Paths   map[string]PathItem `json:"paths"`
}

type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type Server struct {
	URL string `json:"url"`
}

// PathItem holds the HTTP methods for one path.
type PathItem struct {
	Get    *Operation `json:"get,omitempty"`
	Post   *Operation `json:"post,omitempty"`
	Patch  *Operation `json:"patch,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
	Put    *Operation `json:"put,omitempty"`
}

type Operation struct {
	OperationID string      `json:"operationId"`
	Tags        []string    `json:"tags"`
	Summary     string      `json:"summary"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters"`
	RequestBody *RequestBody `json:"requestBody,omitempty"`
	Responses   map[string]Response `json:"responses"`
}

type Parameter struct {
	Name        string      `json:"name"`
	In          string      `json:"in"` // query | path | header
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Schema      *SchemaRef  `json:"schema,omitempty"`
}

type RequestBody struct {
	Description string                     `json:"description"`
	Required    bool                       `json:"required"`
	Content     map[string]MediaTypeObject `json:"content"`
}

type MediaTypeObject struct {
	Schema *SchemaRef `json:"schema,omitempty"`
}

type Response struct {
	Description string                     `json:"description"`
	Content     map[string]MediaTypeObject `json:"content,omitempty"`
}

// SchemaRef is an inline schema or a $ref pointer.
type SchemaRef struct {
	Ref        string            `json:"$ref,omitempty"`
	Type       string            `json:"type,omitempty"`
	Format     string            `json:"format,omitempty"`
	Enum       []any             `json:"enum,omitempty"`
	Items      *SchemaRef        `json:"items,omitempty"`
	Properties map[string]*SchemaRef `json:"properties,omitempty"`
	Required   []string          `json:"required,omitempty"`
	Minimum    *float64          `json:"minimum,omitempty"`
	Maximum    *float64          `json:"maximum,omitempty"`
}

// Op is a flattened, parsed representation of a single API operation.
type Op struct {
	OperationID string
	Method      string
	Path        string
	Tags        []string
	Parameters  []Parameter
	HasBody     bool
}
