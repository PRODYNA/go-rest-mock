package tmpl

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"net/http"
)


type MockResponseWriter struct{}

func (m MockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m MockResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (m MockResponseWriter) WriteHeader(statusCode int) {}

func Test_ConvertTemplate(t *testing.T) {
	r:= &http.Request{}
	r.URL = &url.URL{}
	assert.NoError(t, ConvertTemplate(MockResponseWriter{},"../test/data/backend/sample.tmpl",r))

	assert.Error(t, ConvertTemplate(MockResponseWriter{},"xxx",r))
}
