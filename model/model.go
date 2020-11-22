package model

type MockDefinition struct {
	ID    string `json:"id"`
	Port  string `json:"port"`
	Paths []Path `json:"paths"`
}

type Response struct {
	Status      int               `json:"status"`
	ContentType string            `json:"contentType"`
	Body        map[string]string `json:"body",omitempty`
	// TODO reference to an external file
	BodyRef string            `json:"bodyRef,omitempty"`
	Header  map[string]string `json:"header,omitempty"`
}

type Path struct {
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	ContentType string   `json:"contentType"`
	Response    Response `json:"response,omitempty"`
}
