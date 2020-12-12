package handler

import (
	"github.com/prodyna/go-rest-mock/model"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHandler_NewHandler(t *testing.T) {

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
		Paths: []model.Path{p1, p2, p3},
	}
	NewHandler(&m)
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
}


func Test_getTemplatePath(t *testing.T) {

	p1 := model.Path{
		Method:      "GET",
		Path:        "/test/{test}/user/{id}",
		ContentType: "application/json",
		Response:    model.Response{},
	}

	m := model.MockDefinition{
		Paths: []model.Path{p1},
	}
	h := NewHandler(&m)

	p:= h.getTemplatePath("/test/{test}/user/{id}","GET|application/json")
	assert.Equal(t,&p1,p)

}
