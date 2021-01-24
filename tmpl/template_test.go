package tmpl

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

type MockResponseWriter struct{}

func (m MockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m MockResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (m MockResponseWriter) WriteHeader(int) {}

type MockResponseWriterWithRecording struct {
	header http.Header
	body   []byte
}

func NewMockResponseWriterWithHeader() *MockResponseWriterWithRecording {
	return &MockResponseWriterWithRecording{
		header: http.Header{},
	}
}

func (m MockResponseWriterWithRecording) Header() http.Header {
	return m.header
}

func (m *MockResponseWriterWithRecording) Write(body []byte) (int, error) {
	if m.body == nil {
		m.body = body
	} else {
		m.body = append(m.body, body...)
	}

	return 0, nil
}

func (m MockResponseWriterWithRecording) WriteHeader(int) {}


func Test_ConvertTemplate(t *testing.T) {
	r := &http.Request{}
	r.URL = &url.URL{}
	assert.NoError(t, ConvertTemplate(MockResponseWriter{}, "../test/data/backend/sample.tmpl", r))

	assert.Error(t, ConvertTemplate(MockResponseWriter{}, "xxx", r))
}


func Test_Template(t *testing.T) {
	r := &http.Request{}
	r.URL = &url.URL{}
	m := NewMockResponseWriterWithHeader()
	err :=  ConvertTemplate(m, "../test/data/backend/sample.tmpl", r)
	assert.Nil(t, err)

	assert.Contains(t, string(m.body), "\"length\" : \"4\"")
}
