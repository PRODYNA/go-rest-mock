package model

// MockDefinition is the top level struct for the json config
type MockDefinition struct {
	ID    string `json:"id"`
	Port  string `json:"port"`
	Paths []Path `json:"paths"`
}

// Response definition
type Response struct {
	Status      int                    `json:"status"`
	ContentType string                 `json:"contentType"`
	Body        map[string]interface{} `json:"body"`
	BodyArr     []interface{}          `json:"bodyArr"`
	BodyRef     string                 `json:"bodyRef,omitempty"`
	Header      map[string]string      `json:"header,omitempty"`
}

// Path definition
type Path struct {
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	ContentType string   `json:"contentType"`
	Response    Response `json:"response,omitempty"`
}
