package gofunc

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func Sha1Lower(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1Upper(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}


