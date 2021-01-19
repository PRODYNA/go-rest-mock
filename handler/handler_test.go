package handler

import (
	"encoding/json"
	"github.com/prodyna/go-rest-mock/config"
	"github.com/prodyna/go-rest-mock/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestHandler_NewHandler(t *testing.T) {

	c := &config.Config{}

	p1 := model.Path{
		Method:      "GET",
		Path:        "/test/{test}/user/{id}",
		ContentType: "application/json",
		Response:    model.Response{},
	}
	p2 := model.Path{
		Method:      "POST",
		Path:        "/test/{test}/user/{id}",
		ContentType: "application/json",
		Response:    model.Response{},
	}
	p3 := model.Path{
		Method:      "POST",
		Path:        "_default",
		ContentType: "application/json",
		Response:    model.Response{},
	}
	m := model.MockDefinition{
		Paths:    []model.Path{p1, p2, p3},
		Validate: true,
	}
	NewHandler(&m, c)
}

func Test_validate(t *testing.T) {
	r := http.Request{}
	assert.True(t, validate(&r))

	reader := strings.NewReader("{}")
	r.Header = http.Header{}
	r.Header.Set("content-type", "application/json")
	r.Body = ioutil.NopCloser(reader)
	assert.True(t, validate(&r))

	reader = strings.NewReader("{-}")
	assert.False(t, validate(&r))
}

func Test_isJSONString(t *testing.T) {
	assert.True(t, isJSONString([]byte("{}")))
	assert.False(t, isJSONString([]byte("{-}")))
	assert.False(t, isJSONString([]byte("")))
}

func Test_getContentType(t *testing.T) {

	r := http.Request{}
	assert.Equal(t, "", getContentType(&r))

	r.Header = http.Header{}
	assert.Equal(t, "", getContentType(&r))

	r.Header.Set("content-type", "text/plain")
	assert.Equal(t, "text/plain", getContentType(&r))

	r.Header.Set("content-type", "application/json")
	assert.Equal(t, "application/json", getContentType(&r))
}

func Test_getTemplatePath(t *testing.T) {

	c := &config.Config{}
	p1 := model.Path{
		Method:      "GET",
		Path:        "/test/{test}/user/{id}",
		ContentType: "application/json",
		Response:    model.Response{},
	}

	m := model.MockDefinition{
		Paths:    []model.Path{p1},
		Validate: true,
	}
	h := NewHandler(&m, c)

	p := h.getTemplatePath("/test/{test}/user/{id}", "GET|application/json")
	assert.Equal(t, &p1, p)

}

func Test_hasTemplate(t *testing.T) {

	c := &config.Config{}

	p1 := model.Path{
		Method:      "GET",
		Path:        "/test/{test}/user/{id}",
		ContentType: "application/json",
		Response:    model.Response{},
	}

	m := model.MockDefinition{
		Paths:    []model.Path{p1},
		Validate: true,
	}
	h := NewHandler(&m, c)

	p := h.hasTemplate("", "")
	assert.Equal(t, false, p)

	p = h.hasTemplate("", "GET|application/json")
	assert.Equal(t, false, p)

	p = h.hasTemplate("/test/{test}/user/{id}", "GET|application/json")
	assert.Equal(t, true, p)
}

func Test_getDefault(t *testing.T) {

	c := &config.Config{}

	p1 := model.Path{
		Method:      "",
		Path:        "_default",
		ContentType: "",
		Response:    model.Response{},
	}
	m := model.MockDefinition{
		Paths:    []model.Path{p1},
		Validate: true,
	}
	h := NewHandler(&m, c)
	p := h.getDefault()
	assert.Equal(t, &p1, p)

}

func Test_getStaticPath(t *testing.T) {
	c := &config.Config{}

	p1 := model.Path{
		Method:      "POST",
		Path:        "/api/v1/user/33",
		ContentType: "application/json",
		Response:    model.Response{},
	}
	m := model.MockDefinition{
		Paths:    []model.Path{p1},
		Validate: true,
	}
	h := NewHandler(&m, c)

	p := h.getStaticPath("")
	assert.Nil(t, p)

	p = h.getStaticPath("POST|application/json|/api/v1/user/33")
	assert.Equal(t, &p1, p)
}

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
	m.body = body
	return 0, nil
}

func (m MockResponseWriterWithRecording) WriteHeader(int) {}

func TestHandler_reply(t *testing.T) {
	c := &config.Config{}
	p := model.Path{
		Response: model.Response{
			ContentType: "application/json",
		},
	}
	mrw := NewMockResponseWriterWithHeader()
	reply(mrw, p, c)
	assert.Equal(t, "application/json", mrw.header.Get("Content-Type"))
}

func TestHandler_replyArray(t *testing.T) {
	c := &config.Config{}
	p := model.Path{
		Response: model.Response{
			Body: json.RawMessage("[\"first\",\"second\"]"),
		},
	}
	mrw := NewMockResponseWriterWithHeader()
	reply(mrw, p, c)
	assert.Equal(t, "[\"first\",\"second\"]", string(mrw.body))
}

func TestHandler_ServeHTTP(t *testing.T) {
	c := &config.Config{}

	p1 := model.Path{
		Method:      "POST",
		Path:        "/api/v1/user/33",
		ContentType: "application/json",
		Response:    model.Response{},
	}

	p2 := model.Path{
		Method:      "POST",
		Path:        "/api/v1/book",
		ContentType: "application/json",
		Response:    model.Response{},
	}

	p3 := model.Path{
		Method:      "",
		Path:        "_default",
		ContentType: "",
		Response:    model.Response{},
	}
	m := model.MockDefinition{
		Paths:    []model.Path{p1, p2, p3},
		Validate: true,
	}
	h := NewHandler(&m, c)

	r := http.Request{}
	r.URL = &url.URL{Path: "/api/v1/user/33/XXX"}
	h.ServeHTTP(MockResponseWriter{}, &r)

	r.URL = &url.URL{Path: "/api/v1/user/33"}
	h.ServeHTTP(MockResponseWriter{}, &r)

	r.URL = &url.URL{Path: "/api/v1/user/33"}
	r.Body = ioutil.NopCloser(strings.NewReader("{-}"))
	r.Header = http.Header{}
	r.Header["Content-Type"] = []string{"application/json"}
	h.ServeHTTP(MockResponseWriter{}, &r)

	r.Method = "POST"
	r.URL = &url.URL{Path: "/api/v1/book"}
	r.Header["Content-Type"] = []string{"application/json"}
	r.Body = ioutil.NopCloser(strings.NewReader("{}"))
	h.ServeHTTP(MockResponseWriter{}, &r)

	r.Method = "GET"
	r.URL = &url.URL{Path: "/favicon.ico"}
	h.ServeHTTP(MockResponseWriter{}, &r)
}

func TestHandler_ServeHTTP_Validation(t *testing.T) {
	c := &config.Config{}

	p1 := model.Path{
		Method:      "POST",
		Path:        "/api/v1/book",
		ContentType: "application/json",
		Response:    model.Response{},
	}

	m := model.MockDefinition{
		Paths:    []model.Path{p1},
		Validate: true,
	}
	h := NewHandler(&m, c)

	r := http.Request{}
	r.Header = http.Header{}

	r.Method = "POST"
	r.URL = &url.URL{Path: "/api/v1/book"}
	r.Header["Content-Type"] = []string{"application/json"}
	r.Body = ioutil.NopCloser(strings.NewReader(""))

	mrw := NewMockResponseWriterWithHeader()
	h.ServeHTTP(mrw, &r)

	assert.Equal(t, string(mrw.body), "{ \"error\" : \"Body is invalid\" }")
}

func TestHandler_ServeHTTP_NoValidation(t *testing.T) {
	c := &config.Config{}

	p1 := model.Path{
		Method:      "POST",
		Path:        "/api/v1/book",
		ContentType: "application/json",
		Response:    model.Response{},
	}

	m := model.MockDefinition{
		Paths:    []model.Path{p1},
		Validate: false,
	}
	h := NewHandler(&m, c)

	r := http.Request{}
	r.Header = http.Header{}

	r.Method = "POST"
	r.URL = &url.URL{Path: "/api/v1/book"}
	r.Header["Content-Type"] = []string{"application/json"}
	r.Body = ioutil.NopCloser(strings.NewReader(""))

	mrw := NewMockResponseWriterWithHeader()
	h.ServeHTTP(mrw, &r)

	assert.Equal(t, string(mrw.body), "")
}
