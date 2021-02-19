package tmpl

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var lenFunc = func(in string) int { return len(in) }

var byteArray = func(in string) []byte { return []byte(in) }

var dateIso = func() string { return time.Now().Format("2006-01-02") }

var dateFmt = func(df string) string { return time.Now().Format(df) }

var uuidFunc = func() string { return uuid.New().String() }

var md5Func = func(in string) string {
	hash := md5.Sum([]byte(in))
	return hex.EncodeToString(hash[:])
}

var hexFunc = func(in []byte) string {
	return hex.EncodeToString(in)
}

var stringFunc = func(in []byte) string {
	return string(in)
}

var encodeBase64 = func(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

var decodeBase64 = func(in string) string {
	b, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return "ERROR - not a base64 encoded string"
	}
	return string(b)
}

var decodeBase64Default = func(in,def string) string {
	b, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return def
	}
	return string(b)
}


var sha256Func = func(in string) string  {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(in)))
}
