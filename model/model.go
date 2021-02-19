package model

import "encoding/json"

// MockDefinition is the top level struct for the json config
type MockDefinition struct {
	ID       string `json:"id"`
	Port     string `json:"port"`
	Paths    []Path `json:"paths"`
	Validate bool   `json:"validate"`
}

// Response definition
type Response struct {
	Status      int               `json:"status"`
	ContentType string            `json:"contentType"`
	Body        json.RawMessage   `json:"body"`
	BodyRef     string            `json:"bodyRef,omitempty"`
	TemplateRef string            `json:"templateRef,omitempty"`
	Header      map[string]string `json:"header,omitempty"`
}

// Path definition
type Path struct {
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	ContentType string   `json:"contentType"`
	Response    Response `json:"response,omitempty"`
}
