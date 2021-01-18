package handler

import (
	"encoding/json"
	"fmt"
	"github.com/prodyna/go-rest-mock/config"
	"github.com/prodyna/go-rest-mock/model"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	defaultKey  = "||_default"
	contentType = "Content-Type"
	appJson     = "application/json"
)

// Handler is the http handler
type Handler struct {
	// static paths
	staticMap map[string]model.Path

	// paths with templates. first key is method and content type, second key is path
	templateMap map[string]map[string]model.Path

	config *config.Config
}

// NewHandler creates a handler for the configuration.
func NewHandler(md *model.MockDefinition, conf *config.Config) *Handler {

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

	return &Handler{staticMap: staticMap, templateMap: templateMap, config: conf}
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

// ServeHTTP is the handler for one mock configuration.
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// in case you request it with a browser :-)
	if req.URL.Path == "/favicon.ico" && req.Method == "GET" {
		w.WriteHeader(204)
		return
	}

	if !validate(req) {
		w.Header().Set(contentType, appJson)
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
		reply(w, *staticPath, h.config)
		return
	}

	templatePath := h.getTemplatePath(reqPath, templateKey)
	if templatePath != nil {
		reply(w, *templatePath, h.config)
		return
	}

	defaultPath := h.getDefault()
	if defaultPath != nil {
		reply(w, *defaultPath, h.config)
		return
	}

	w.Header().Set(contentType, appJson)
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
		return false
	}
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

// Replies with the configured data.
func reply(w http.ResponseWriter, path model.Path, cfg *config.Config) {

	status := path.Response.Status
	respContentType := path.Response.ContentType
	var respBody []byte
	if path.Response.BodyRef != "" {
		respBody, _ = ioutil.ReadFile(cfg.Path + "/" + path.Response.BodyRef)
	} else {
		body := path.Response.Body
		respBody, _ = json.Marshal(body)
	}

	w.Header().Set(contentType, respContentType)

	for key, header := range path.Response.Header {
		w.Header().Set(key, header)
	}

	w.WriteHeader(status)
	w.Write(respBody)
}

// Save method for getting the content type
func getContentType(req *http.Request) string {
	if req.Header == nil {
		return ""
	} else if len(req.Header) == 0 {
		return ""
	} else {
		h := req.Header[contentType]
		if h == nil {
			return ""
		}
		return h[0]
	}
}
