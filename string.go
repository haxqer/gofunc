package gofunc

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
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

func Base64EncodeByte(b []byte) []byte {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(base64Text, b)
	return base64Text
}

func Base64DecodeByte(b []byte) ([]byte, error) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	l, err := base64.StdEncoding.Decode(base64Text, b)
	if err != nil {
		return nil, err
	}
	return base64Text[:l], nil
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

func HexEncode(s string) string {
	return hex.EncodeToString([]byte(s))
}

func HexEncodeByte(b []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	return dst
}

func HexDecode(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

func HexDecodeByte(b []byte) ([]byte, error) {
	n, err := hex.Decode(b, b)
	return b[:n], err
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}



