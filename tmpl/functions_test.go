package tmpl

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_sha256Func(t *testing.T) {
	s := sha256Func("data")
	assert.Equal(t, "3a6eb0790f39ac87c94f3856b2dd2c5d110e6811602261a9a923d3bb23adc8b7", s)
}

func Test_stringFunc(t *testing.T) {
	s := stringFunc([]byte("data"))
	assert.Equal(t, "data", s)
}

func Test_md5Func(t *testing.T) {
	s := md5Func(("data"))
	assert.Equal(t, "8d777f385d3dfec8815d20f7496026dc", s)
}

func Test_uuidFunc(t *testing.T) {
	s := uuidFunc()
	assert.Equal(t, 36, len(s))
	assert.Regexp(t, "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", s)
}

func Test_dateIso(t *testing.T) {
	s := dateIso()
	assert.Equal(t, time.Now().Format("2006-01-02"), s)
}

func Test_dateFmt(t *testing.T) {
	s := dateFmt("2006-01-02 XXX")
	assert.Equal(t, time.Now().Format("2006-01-02 XXX"), s)
}

func Test_lenFunc(t *testing.T) {
	l := lenFunc("1234")
	assert.Equal(t, 4, l)
}

func Test_hexFunc(t *testing.T) {
	h := hexFunc([]byte( "abc" ))
	assert.Equal(t, "616263", h)
}

func Test_byteArray(t *testing.T) {
	h := byteArray("abc")
	assert.Equal(t, []byte("abc"), h)
}

func Test_encodeBase64(t *testing.T) {
	b := encodeBase64("abc")
	assert.Equal(t, "YWJj", b)
}

func Test_decodeBase64(t *testing.T) {
	b := decodeBase64("YWJj")
	assert.Equal(t, "abc", b)
}

func Test_decodeBase64_Error(t *testing.T) {
	b := decodeBase64("YWJj2")
	assert.Equal(t, "ERROR - not a base64 encoded string", b)
}

func Test_decodeBase64Default(t *testing.T) {
	b := decodeBase64Default("YWJj2", "")
	assert.Equal(t, "", b)

	b = decodeBase64Default("YWJj", "")
	assert.Equal(t, "abc", b)
}
