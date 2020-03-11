package gofunc

import (
	"bytes"
	"github.com/pquerna/ffjson/ffjson"
)

func MarshalFalseEscapeHTML(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := ffjson.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	r := buffer.Bytes()
	if len(r) > 0 {
		r = r[:len(r)-1]
	}
	return r, err
}
