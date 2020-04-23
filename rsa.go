package gofunc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)


func RsaEncodeWithPublic(input []byte, rsaKey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, rsaKey, input)
}

func RsaDecodeWithPrivate(input []byte, rsaKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, rsaKey, input)
}

func RsaDecode(input, key []byte) ([]byte, error) {
	rsaKey, err := ParsePrivateKeyFromKeyBytes(key)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, rsaKey, input)
}

func RsaEncode(input, key []byte) ([]byte, error) {
	rsaKey, err := ParsePublicKeyFromKeyBytes(key)
	if err != nil {
		return nil, err
	}

	return rsa.EncryptPKCS1v15(rand.Reader, rsaKey, input)
}

func ParsePrivateKeyFromKeyBytes(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, fmt.Errorf("pem.Decode error: %+v ", key)
	}
	enc := x509.IsEncryptedPEMBlock(block)
	var err error
	b := block.Bytes
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}

	rsaKey, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, err
	}

	return rsaKey, nil
}

func ParsePublicKeyFromKeyBytes(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, fmt.Errorf("pem.Decode error: %+v ", key)
	}
	enc := x509.IsEncryptedPEMBlock(block)
	var err error
	b := block.Bytes
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}

	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return nil, err
	}

	return rsaKey, nil
}


