package gofunc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"reflect"
	"testing"
)

func TestAesCBCEncrypt(t *testing.T) {
	type args struct {
		plainText   []byte
		key         []byte
		paddingFunc func([]byte, int) []byte
	}

	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "testCase01",
			args: args{
				plainText:   []byte("exampleplaintext12"),
				key:         key,
				paddingFunc: PKCS5Padding,
			},
			want: []byte("exampleplaintext12"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesCBCEncrypt(tt.args.plainText, tt.args.key, tt.args.paddingFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCBCEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			want2, err := AesCBCDecrypt(got, key,PKCS5UnPadding)
			if err != nil {
				t.Errorf("AesCBCDecrypt() error = %v", err)
				return
			}

			if !reflect.DeepEqual(want2, tt.want) {
				t.Errorf("AesCBCEncrypt() got = %x, want %x", got, tt.want)
			}
		})
	}
}

func TestExampleNewCBCEncrypter(t *testing.T) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("exampleplaintext")

	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	//fmt.Printf("%x\n", ciphertext)
}

func TestExampleNewCBCDecrypter(t *testing.T) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString("5f14af7a8d9f0e28c58c595f1c4c56fbfeab5e157ea9bd54d23950a4f4cbac7f52a3c602f9bd24bc7f5f4647904a355b")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertext, ciphertext)
	//fmt.Printf("%s\n", ciphertext)
	//fmt.Printf("%+v\n", ciphertext)
}

func TestAesCBCDecrypt(t *testing.T) {
	type args struct {
		cipherText    []byte
		key           []byte
		unPaddingFunc func([]byte) ([]byte, error)
	}
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	cipherText, _ := hex.DecodeString("d5946b8e9461d437a2fbd6828398cd323d04769b535eae18c6c1944dbaff91bc845c546b44b27024705092cc55d1d085")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "testCase01",
			args: args{
				cipherText:   cipherText,
				key:         key,
				unPaddingFunc: PKCS5UnPadding,
			},
			want: []byte("exampleplaintext12"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesCBCDecrypt(tt.args.cipherText, tt.args.key, tt.args.unPaddingFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCBCDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesCBCDecrypt() got = %s, want %v", got, tt.want)
			}
		})
	}
}