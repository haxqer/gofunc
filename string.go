package gofunc

import (
	"encoding/base64"
	"math/rand"
)

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

func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

