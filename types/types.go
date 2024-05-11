package types

// 文档地址
// https://openapi.apifox.cn/

type OpenAPIJson struct {
	Openapi    string                `json:"openapi,omitempty"`
	Servers    []*Servers            `json:"servers,omitempty"`
	Info       *Info                 `json:"info,omitempty"`
	Tags       []*Tag                `json:"tags,omitempty"`
	Security   []map[string][]string `json:"security"`
	Paths      map[string]*PathItem  `json:"paths,omitempty"`
	Components *Components           `json:"components,omitempty"`
}

type Info struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
}

type Servers struct {
	Url         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

type Tag struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Security struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Name        string `json:"name"`
	In          string `json:"in"`
}

type PathItem struct {
	Ref         string        `json:"$ref,omitempty"`
	Summary     string        `json:"summary,omitempty"`
	Description string        `json:"description,omitempty"`
	Parameters  []*Parameters `json:"parameters,omitempty"`
	Get         *Operation    `json:"get,omitempty"`
	Put         *Operation    `json:"put,omitempty"`
	Post        *Operation    `json:"post,omitempty"`
	Delete      *Operation    `json:"delete,omitempty"`
	Options     *Operation    `json:"options,omitempty"`
	Head        *Operation    `json:"head,omitempty"`
	Patch       *Operation    `json:"patch,omitempty"`
	Trace       *Operation    `json:"trace,omitempty"`
	Servers     []*Servers    `json:"servers,omitempty"`
}

type Operation struct {
	Tags        []string             `json:"tags,omitempty"`
	Summary     string               `json:"summary,omitempty"`
	Description string               `json:"description,omitempty"`
	OperationId string               `json:"operationId,omitempty"`
	Parameters  []*Parameters        `json:"parameters,omitempty"`  // Parameters or Reference
	RequestBody *RequestBody         `json:"requestBody,omitempty"` // RequestBody or Reference
	Responses   map[string]*Response `json:"responses,omitempty"`
	Security    map[string][]string  `json:"security,omitempty"`
}

type Parameters struct {
	Name        string `json:"name,omitempty"`
	In          string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

type RequestBody struct {
	Description string                `json:"description,omitempty"`
	Required    bool                  `json:"required,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty"`
}

type Response struct {
	Description string                `json:"description,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty"` //
}

type Reference struct {
	Ref         string `json:"$ref,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}

type MediaType struct {
	Schema   *Schema              `json:"schema,omitempty"`
	Example  any                  `json:"example,omitempty"`
	Encoding map[string]*Encoding `json:"encoding,omitempty"`
}

type Encoding struct {
	ContentType string `json:"contentType,omitempty"`
}

type Schema struct {
	Ref                  string             `json:"$ref,omitempty"`
	Description          string             `json:"description,omitempty"`
	Type                 string             `json:"type,omitempty"`
	Title                string             `json:"title,omitempty"`
	Format               string             `json:"format,omitempty"`
	Items                *Schema            `json:"items,omitempty"`
	Required             []string           `json:"required,omitempty"`
	Properties           map[string]*Schema `json:"properties,omitempty"`
	AdditionalProperties *Schema            `json:"additionalProperties,omitempty"`
}

type Components struct {
	Schemas         map[string]*Schema   `json:"schemas,omitempty"`
	SecuritySchemes map[string]*Security `json:"securitySchemes,omitempty"`
}
