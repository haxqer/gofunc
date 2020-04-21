package gofunc

import (
	"encoding/base64"
	"math/rand"
	"net/url"
	"regexp"
)
// generate random string by array of rune.
func RandStringRunes(n int, strArr []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = strArr[rand.Intn(len(strArr))]
	}
	return string(b)
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func Base64DecodeByte(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// for RFC 3986
// reference: https://tools.ietf.org/html/rfc3986
func RFCEncode(s string) string {
	return url.QueryEscape(s)
}

var RFCRegexp = regexp.MustCompile(`^[0-9a-zA-Z\-_.~+%]*$`)
// for RFC 3986
// check RFC character
func RFCCheck(s string) bool {
	return RFCRegexp.MatchString(s)
}
