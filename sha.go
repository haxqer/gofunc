package gofunc

import (
	"crypto"
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

func Sha256Lower(s string) string  {
	h := crypto.SHA256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

