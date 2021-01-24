package tmpl

import (
	"io/ioutil"
	"net/http"
	"text/template"
)

type TemplateHandler struct {
	AbsCount int64

	// count()
	// everyCount(int)
	// countAbs()
	// reqBodyLength()
	// reqBodyRegex()
	// regBodyJsonPath()
	// randomInt
	// randomFloat
	// randomString(length)
	// reqMethod()
	// reqPath()
	// reqTemplateParam(tpl)
	// Time()
	// Time(fmt)
	// DayOfWeek()
	// Year()
	// Month()
	// Replace(x,y)

}

// ConvertTemplate loads a template and executes it with the http request data
func ConvertTemplate(w http.ResponseWriter, path string, r *http.Request) error {

	templateFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	tmp := template.Must(template.New("").Funcs(functions()).Parse(string(templateFile)))

	err = tmp.Execute(w, r)

	return err
}

func functions() template.FuncMap {
	return template.FuncMap{
		"DateIso":      dateIso,
		"DateFmt":      dateFmt,
		"Uuid":         uuidFunc,
		"MD5":          md5Func,
		"Hex":          hexFunc,
		"String":       stringFunc,
		"EncodeBase64": encodeBase64,
		"DecodeBase64": decodeBase64,
		"ByteArray":    byteArray,
		"Length":       lenFunc,
		"Sha256":       sha256Func,
	}
}
