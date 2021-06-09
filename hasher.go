package gofunc

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func Md5Lower(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5Upper(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func HmacSha1(s, key string) []byte {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(s))
	return h.Sum(nil)
}

func HmacSha1HexString(s, key string) string {
	return hex.EncodeToString(HmacSha1(s, key))
}

func HmacSha1Base64(s, key string) string {
	return base64.StdEncoding.EncodeToString(HmacSha1(s, key))
}


func HmacSha256(s, key string) []byte {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(s))
	return h.Sum(nil)
}

func HmacSha256HexString(s, key string) string {
	return string(HmacSha1(s, key))
}
