package gofunc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func AesCBCEncrypt(plainText []byte, key []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainText = paddingFunc(plainText, aes.BlockSize)

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return cipherText, nil
}

func AesCBCDecrypt(cipherText []byte, key []byte, unPaddingFunc func([]byte) ([]byte, error)) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	plainText, err := unPaddingFunc(cipherText)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
