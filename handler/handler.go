package handler

import (
	"encoding/json"
	"fmt"
	"github.com/prodyna/go-rest-mock/model"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	defaultKey  = "||_default"
	ContentType = "Content-Type"
	AppJson     = "application/json"
)

// The Http handler
type Handler struct {
	// static paths
	staticMap map[string]model.Path

	// paths with templates. first key is method and content type, second key is path
	templateMap map[string]map[string]model.Path
}

// Creates a handler for the configuration.
func NewHandler(md *model.MockDefinition) *Handler {

	staticMap := make(map[string]model.Path)
	templateMap := make(map[string]map[string]model.Path)

	for _, path := range md.Paths {
		if strings.Contains(path.Path, "{") {
			key := strings.Join([]string{path.Method, path.ContentType}, "|")
			if _, ok := templateMap[key]; ok {
				templateMap[key][path.Path] = path
			} else {
				templateMap[key] = make(map[string]model.Path)
				templateMap[key][path.Path] = path
			}
		} else {
			key := strings.Join([]string{path.Method, path.ContentType, path.Path}, "|")
			staticMap[key] = path
		}
	}

	return &Handler{staticMap: staticMap, templateMap: templateMap}
}

func (h *Handler) getStaticPath(key string) *model.Path {
	if _, ok := h.staticMap[key]; ok {
		p := h.staticMap[key]
		return &p
	}
	return nil
}

func (h *Handler) getDefault() *model.Path {
	if _, ok := h.staticMap[defaultKey]; ok {
		p := h.staticMap[defaultKey]
		return &p
	}
	return nil
}

// A handler for one mock configuration.
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// in case you request it with a browser :-)
	if req.URL.Path == "/favicon.ico" && req.Method == "GET" {
		w.WriteHeader(204)
		return
	}

	if !validate(req) {
		w.Header().Set(ContentType, AppJson)
		w.WriteHeader(400)
		fmt.Fprintf(w, "{ \"error\" : \"Body is invalid\" }")
		return
	}

	reqPath := req.URL.Path
	method := req.Method
	contentType := getContentType(req)

	key := strings.Join([]string{method, contentType, reqPath}, "|")
	templateKey := strings.Join([]string{method, contentType}, "|")

	staticPath := h.getStaticPath(key)
	if staticPath != nil {
		reply(w, *staticPath)
		return
	}

	templatePath := h.getTemplatePath(reqPath, templateKey)
	if templatePath != nil {
		reply(w, *templatePath)
		return
	}

	defaultPath := h.getDefault()
	if defaultPath != nil {
		reply(w, *defaultPath)
		return
	}

	w.Header().Set(ContentType, AppJson)
	w.WriteHeader(404)
	fmt.Fprintf(w, "{\"error\" : \"no mapping for "+req.URL.Path+"\"}")

}

func (h *Handler) hasTemplate(reqPath string, templateKey string) bool {

	if _, ok := h.templateMap[templateKey]; ok {

		pmap := h.templateMap[templateKey]

		for k := range pmap {
			if Match(reqPath, k) {
				return true
			}
		}
	}

	return false
}

func (h *Handler) getTemplatePath(reqPath string, templateKey string) *model.Path {

	if _, ok := h.templateMap[templateKey]; ok {

		pMap := h.templateMap[templateKey]

		for k, v := range pMap {
			if Match(reqPath, k) {
				return &v
			}
		}
	}

	return nil
}

func validate(req *http.Request) bool {

	if req.Method != "GET" && getContentType(req) == "application/json" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return false
		}
		return isJSONString(body)
	}
	return true
}

func isJSONString(s []byte) bool {
	if len(s) == 0 {
		return true
	}
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

/**
 * Replies with the configured data.
 */
func reply(w http.ResponseWriter, path model.Path) {
	body := path.Response.Body
	status := path.Response.Status
	contentType := path.Response.ContentType
	jsonString, _ := json.Marshal(body)

	w.Header().Set(ContentType, contentType)

	for key, header := range path.Response.Header {
		w.Header().Set(key, header)
	}

	w.WriteHeader(status)
	w.Write(jsonString)

}

/**
 * Save method for getting the content type
 */
func getContentType(req *http.Request) string {
	if req.Header == nil {
		return ""
	} else if len(req.Header) == 0 {
		return ""
	} else {
		h := req.Header[ContentType]
		if h == nil {
			return ""
		}
		return h[0]
	}
}
