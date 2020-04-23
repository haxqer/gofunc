package gofunc

import "fmt"

type Response struct {
	EK []byte
	ED []byte
}
var private1KeyBytes = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCyBUxPA9X3tdJSsgNi0/CctBsBbvJEUgt5FWbbG4nTp2r7XK2T
Vk/YBZ6jq7kiA1rzJ/xmJTdpUdVUhPfX1DN/7iapb5K3z/NhIO/jsJgGO+YtgK5I
WEcjGwElPZnOsMk6iNZAWCGa7EyA0FHkul7w4eFOjC+RGqbKfsl306EUtwIDAQAB
AoGBAIMp8ip5ugoUVk4FyQbk/3CGJyusMiZyiO+C/FDN/oQK44EmrOFVA+k3YsZW
/UX5UOa9fHNKUoRv/g2TFwVX3UTVuri4h4x6zPfCOv3O1TUVaJzgCGa8bc/sNRu4
AlYZZdeKNxQbURbht/kbyu11HjyPqPP+ZXcwoalCiZQmtI/xAkEA7VbfGeQuH8Go
4SBv3EbQvFhJBWjjxmjIq3VQZD7HqdZMTqYnqaq3bVvqHWizrxyku2BbgaFbV5MJ
ZhbeZOF/bwJBAMAEeGbjs9xkPcfMcOCVypZpRcoi/0BMzj6Yl8YKmzxY6cNVIyEB
UukklSxZlbt8piFaZHZo1e85g1pDiX01GzkCQBvTx6y9eDr49dgPeY4WL3slzsn3
ll05A+42fwqB4d8j5SaDjLrz7TXBRR3VnNu3PAlMLu5wAMmvz7ZMkB674bkCQQCE
ZlCy+Uz6qW/kBXbLlN2Eyv/hOjJwnsUTalo0pvmVKeW910WKq4QE2EG3u+m/xloy
40YkU3M4KZsFsU3rNKQZAkAdJJEhk3+LwoCSD/hneL7aCtCAmmmWOuHTWjQAD1Zo
OEJTQoc9/wQ767QpGcbqdwF+KjaEVTqn5Wx+9UkMWhyy
-----END RSA PRIVATE KEY-----
`)

var public1KeyBytes = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCyBUxPA9X3tdJSsgNi0/CctBsB
bvJEUgt5FWbbG4nTp2r7XK2TVk/YBZ6jq7kiA1rzJ/xmJTdpUdVUhPfX1DN/7iap
b5K3z/NhIO/jsJgGO+YtgK5IWEcjGwElPZnOsMk6iNZAWCGa7EyA0FHkul7w4eFO
jC+RGqbKfsl306EUtwIDAQAB
-----END PUBLIC KEY-----
`)

func Pack(b []byte) (*Response, error) {
	// 随机生成一个32位的 key
	key, err := GenRandomEK()
	if err != nil {
		return nil, err
	}
	fmt.Printf("key: %X \n", key)
	// rsa 公钥加密 key, 生成加密后的 ek
	ek, err := RsaEncode(key, public1KeyBytes)
	if err != nil {
		return nil, err
	}
	// 将加密后的 ek 进行 Base64Encode
	bek := Base64EncodeByte(ek)
	fmt.Printf("bek: %s \n", bek)

	fmt.Printf("ek: %X \n", ek)
	// 使用上一步随机生成的 key, 随机生成一个 IV (这里已经封装在 AesCBCEncrypt中), 生成 ed
	ed, err := AesCBCEncrypt(b, key, PKCS5Padding)
	if err != nil {
		 return nil, err
	}
	fmt.Printf("ed: %X \n", ed)
	// 将加密后的 ed 进行 Base64Encode
	bed := Base64EncodeByte(ed)

	fmt.Printf("bed: %s \n", bed)
	return &Response{
		EK: bek,
		ED: bed,
	}, nil
}

func Unpack(r *Response) ([]byte, error) {
	ek, err := Base64DecodeByte(r.EK)
	if err != nil {
		return nil, err
	}
	fmt.Printf("ek: %X \n", ek)
	key, err := RsaDecode(ek, private1KeyBytes)
	if err != nil {
		return nil, err
	}
	fmt.Printf("key: %X \n", key)

	ed, err := Base64DecodeByte(r.ED)
	if err != nil {
		return nil, err
	}
	b, err := AesCBCDecrypt(ed, key, PKCS5UnPadding)
	if err != nil {
		return nil, err
	}
	fmt.Printf("b: %s \n", b)

	return b, nil
}



