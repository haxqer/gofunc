package gofunc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

func Aes128CBCEncryptBase64Encode(key []byte, iv []byte, src []byte) ([]byte, error) {
	result, err := Aes128CBCEncrypt(src, key, iv)
	if err != nil {
		return nil, err
	}
	return []byte(base64.StdEncoding.EncodeToString(result)), nil
}

func Aes128CBCDecryptBase64Decode(key []byte, iv []byte, src []byte) ([]byte, error) {
	result, err := base64.StdEncoding.DecodeString(string(src))
	if err != nil {
		return nil, err
	}
	origData, err := Aes128CBCDecrypt(result, key, iv)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

func Aes128CBCEncrypt(origData, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, errors.New(fmt.Sprintf("Aes128CBCDecrypt error key : %+v", key))
	}
	if IV != nil && len(IV) != 16 {
		return nil, errors.New(fmt.Sprintf("Aes128CBCDecrypt error IV : %+v", IV))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, IV[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func Aes128CBCDecrypt(crypted, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, errors.New(fmt.Sprintf("Aes128CBCDecrypt error key : %+v", key))
	}
	if IV != nil && len(IV) != 16 {
		return nil, errors.New(fmt.Sprintf("Aes128CBCDecrypt error IV : %+v", IV))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(crypted)%aes.BlockSize != 0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, IV[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData, err = PKCS5UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

func AesCBCEncrypt(plainText []byte, key []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainText = paddingFunc(plainText, aes.BlockSize)

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
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
