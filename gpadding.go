package gofunc

import (
	"bytes"
	"errors"
)

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS5UnPadding(plainText []byte) ([]byte, error) {
	if len(plainText) == 0 {
		return nil, errors.New("unPadding error: length of plainText is zero")
	}
	length := len(plainText)
	unPadding := int(plainText[length - 1])
	if length < unPadding {
		return nil, errors.New("unPadding error: unPadding bigger than length")
	}
	return plainText[:(length - unPadding)], nil
}
