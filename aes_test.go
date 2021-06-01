package gofunc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestGenRandomIV(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{name:"testCase01", want: nil, wantErr: false},
		{name:"testCase02", want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenRandomIV()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenRandomIV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != 16 {
				t.Errorf("GenRandomIV() error = length is not 32")
				return
			}
			//fmt.Printf("%X \n", got)
		})
	}
}

func BenchmarkGenRandomIV(b *testing.B) {
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := GenRandomIV()
		if err != nil {
			b.Errorf("GenRandomIV() err %+v = got = %s", err, got)
		}
		if len(got) != 16 {
			b.Errorf("GenRandomIV() error = length is not 32")
			return
		}
	}
}

func TestGenRandomEK(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{name:"testCase01", want: nil, wantErr: false},
		{name:"testCase02", want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenRandomEK()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenRandomEK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != 32 {
				t.Errorf("GenRandomEK() error = length is not 32")
				return
			}
			//fmt.Printf("%X \n", got)
		})
	}
}

func BenchmarkGenRandomEK(b *testing.B) {
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := GenRandomEK()
		if err != nil {
			b.Errorf("GenRandomEK() err %+v = got = %s", err, got)
		}
		if len(got) != 32 {
			b.Errorf("GenRandomEK() error = length is not 32")
			return
		}
	}
}

func TestAesCBCEncrypt(t *testing.T) {
	type args struct {
		plainText   []byte
		key         []byte
		paddingFunc func([]byte, int) []byte
	}

	key, _ := hex.DecodeString("DDF11F26FC4321EB0C2B83C9D1B2120B64716E4FF28FB3D0E8F4FF73B381A4CC")
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
			want:    []byte("exampleplaintext12"),
			wantErr: false,
		},
		{
			name: "testCase02",
			args: args{
				plainText:   []byte("exampleplainte"),
				key:         key,
				paddingFunc: PKCS5Padding,
			},
			want:    []byte("exampleplainte"),
			wantErr: false,
		},
		{
			name: "testCase03",
			args: args{
				plainText:   []byte("0123456789012345678901234567890123456789"),
				key:         key,
				paddingFunc: PKCS5Padding,
			},
			want:    []byte("0123456789012345678901234567890123456789"),
			wantErr: false,
		},
		{
			name: "testCase03",
			args: args{
				plainText:   []byte("123"),
				key:         key,
				paddingFunc: PKCS5Padding,
			},
			want:    []byte("123"),
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
			fmt.Printf("%X \n", got)

			want2, err := AesCBCDecrypt(got, key, PKCS5UnPadding)
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
				cipherText:    cipherText,
				key:           key,
				unPaddingFunc: PKCS5UnPadding,
			},
			want:    []byte("exampleplaintext12"),
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

func TestAes128CBCEncryptBase64Encode(t *testing.T) {
	type args struct {
		key []byte
		iv  []byte
		src []byte
	}

	key := []byte("1234567812345678")
	iv := []byte("0000000000000000")

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "testCase01",
			args: args{
				key: key,
				iv: iv,
				src: []byte("我这是测试加密abd134{}[]\\()=-=+_!@$%@#$%^^&&*(_)+~`<>\\?,,./;';:\"\":^^"),
			},
			want: []byte("ehHf5AOFj5vQib3+iU9Ake/qRkCoNGPy2JZ+/f6MONv7Cmb+sSNK3x1kHOmX5amQUEww7a/EaR3ikvtM/u+DFX/8WBVv1L6tQ1gJtp4Emjs="),
			wantErr: false,
		},
		{
			name: "testCase02",
			args: args{
				key: key,
				iv: iv,
				src: []byte("absdefffsss"),
			},
			want: []byte("5EHbdxRjM3ELEkNWBMhx1A=="),
			wantErr: false,
		},
		{
			name: "testCase03",
			args: args{
				key: key,
				iv: iv,
				src: []byte("1234567887654321"),
			},
			want: []byte("t9M3j1tQPRqhAz4Elx8Wpwc5p8YIgRvchoNEFgmTj4U="),
			wantErr: false,
		},
		{
			name: "testCase04",
			args: args{
				key: key,
				iv: iv,
				src: []byte("12345678876543211234567887654321"),
			},
			want: []byte("t9M3j1tQPRqhAz4Elx8Wp8LsxFMY+FhwQp8pUtmEW1fnbzRAJ99SjJ9av9nDnXBq"),
			wantErr: false,
		},
		{
			name: "testCase04",
			args: args{
				key: key,
				iv: iv,
				src: []byte("12345678876543211234567887654321b"),
			},
			want: []byte("t9M3j1tQPRqhAz4Elx8Wp8LsxFMY+FhwQp8pUtmEW1c2rt/G0JFSBi+8nJStGi0b"),
			wantErr: false,
		},
		{
			name: "testCase05",
			args: args{
				key: key,
				iv: iv,
				src: []byte("12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b"),
			},
			want: []byte("t9M3j1tQPRqhAz4Elx8Wp8LsxFMY+FhwQp8pUtmEW1czJYuXQjai3JM5tzc4u+nR8Wi/pDQFHtQDruYFNsRZhhKTgjG6oZQUObfkkmZrRphUYzDzpsS2DM0H+1deGiOvEhdJuJWc2WynKJwYFbOpK+Ibf8H9xnl+6Sy3ycSMpGcweHBAqixfvAL3Z+HK/zwtBjFLd3ku9VoJ2wvrdL6fg4eaIfUXBp0KST9E2E46H9kQIUi54vIopjnhzepxuoHw9aP0t2/hNQWLnl8GUYfs1zVlddsiPq7fH3TN0w+LmSicYLCtT8P3Z93hIJShqaurq0OpZxxeJpDD9bcRSM3CSrwhuAG1IOHXR/oXBhkfzxBbmx108aY0llSAwvm4Au982yytzyb0c2CpsLhUEf4tDCsZ/nH8wuqf5K28UZs1NIczKdQ2If+BGuLmTp2pJB9qZQMT2+9lk4UlQA3gEx63OrteKqFop8m7NSCCMmWlIkHr2ro1VIbY2DoU6cJTR1PAifl1fM7OssuhpoZHgsYNVKZPQXaZ0sH5wH+FzKGsG8o9xoWnwbiXUE79+VjlOjVf4euE0Ud2WpwmM/9YbGSQLjrJ7M+nnpVIAJ2bd5Z570gmIuzaXjGW5h3mqpvHk93/T0Rs71P6+FsnDeLDYIakLH2mN9p0Y4yGJTGyqI9OCnOR9FWiAf1WjMKZr2BUIllVzgnhsiYkBVTXQ9qUgjHzXw=="),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Aes128CBCEncryptBase64Encode(tt.args.key, tt.args.iv, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Aes128CBCEncryptBase64Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aes128CBCEncryptBase64Encode() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestAes128CBCDecryptBase64Decode(t *testing.T) {
	type args struct {
		key []byte
		iv  []byte
		src []byte
	}
	key := []byte("1234567812345678")
	iv := []byte("0000000000000000")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "testCase01",
			args: args{
				key: key,
				iv: iv,
				src: []byte("ehHf5AOFj5vQib3+iU9Ake/qRkCoNGPy2JZ+/f6MONv7Cmb+sSNK3x1kHOmX5amQUEww7a/EaR3ikvtM/u+DFX/8WBVv1L6tQ1gJtp4Emjs="),
			},
			want: []byte("我这是测试加密abd134{}[]\\()=-=+_!@$%@#$%^^&&*(_)+~`<>\\?,,./;';:\"\":^^"),
			wantErr: false,
		},
		{
			name: "testCase02",
			args: args{
				key: key,
				iv: iv,
				src: []byte("5EHbdxRjM3ELEkNWBMhx1A=="),
			},
			want: []byte("absdefffsss"),

			wantErr: false,
		},
		{
			name: "testCase03",
			args: args{
				key: key,
				iv: iv,
				src: []byte("t9M3j1tQPRqhAz4Elx8Wpwc5p8YIgRvchoNEFgmTj4U="),
			},
			want: []byte("1234567887654321"),
			wantErr: false,
		},
		{
			name: "testCase04",
			args: args{
				key: key,
				iv: iv,
				src: []byte("t9M3j1tQPRqhAz4Elx8Wp8LsxFMY+FhwQp8pUtmEW1fnbzRAJ99SjJ9av9nDnXBq"),
			},
			want: []byte("12345678876543211234567887654321"),
			wantErr: false,
		},
		{
			name: "testCase04",
			args: args{
				key: key,
				iv: iv,
				src: []byte("t9M3j1tQPRqhAz4Elx8Wp8LsxFMY+FhwQp8pUtmEW1c2rt/G0JFSBi+8nJStGi0b"),
			},
			want: []byte("12345678876543211234567887654321b"),
			wantErr: false,
		},
		{
			name: "testCase05",
			args: args{
				key: key,
				iv: iv,
				src: []byte("t9M3j1tQPRqhAz4Elx8Wp8LsxFMY+FhwQp8pUtmEW1czJYuXQjai3JM5tzc4u+nR8Wi/pDQFHtQDruYFNsRZhhKTgjG6oZQUObfkkmZrRphUYzDzpsS2DM0H+1deGiOvEhdJuJWc2WynKJwYFbOpK+Ibf8H9xnl+6Sy3ycSMpGcweHBAqixfvAL3Z+HK/zwtBjFLd3ku9VoJ2wvrdL6fg4eaIfUXBp0KST9E2E46H9kQIUi54vIopjnhzepxuoHw9aP0t2/hNQWLnl8GUYfs1zVlddsiPq7fH3TN0w+LmSicYLCtT8P3Z93hIJShqaurq0OpZxxeJpDD9bcRSM3CSrwhuAG1IOHXR/oXBhkfzxBbmx108aY0llSAwvm4Au982yytzyb0c2CpsLhUEf4tDCsZ/nH8wuqf5K28UZs1NIczKdQ2If+BGuLmTp2pJB9qZQMT2+9lk4UlQA3gEx63OrteKqFop8m7NSCCMmWlIkHr2ro1VIbY2DoU6cJTR1PAifl1fM7OssuhpoZHgsYNVKZPQXaZ0sH5wH+FzKGsG8o9xoWnwbiXUE79+VjlOjVf4euE0Ud2WpwmM/9YbGSQLjrJ7M+nnpVIAJ2bd5Z570gmIuzaXjGW5h3mqpvHk93/T0Rs71P6+FsnDeLDYIakLH2mN9p0Y4yGJTGyqI9OCnOR9FWiAf1WjMKZr2BUIllVzgnhsiYkBVTXQ9qUgjHzXw=="),
			},
			want: []byte("12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b"),
			wantErr: false,
		},
		{
			name: "testCase06",
			args: args{
				key: key,
				iv: iv,
				src: []byte("12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b12345678876543211234567887654321b"),
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Aes128CBCDecryptBase64Decode(tt.args.key, tt.args.iv, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Aes128CBCDecryptBase64Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aes128CBCDecryptBase64Decode() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestAesGCMNoPaddingDecryptBase64Decode(t *testing.T) {
	type args struct {
		cipherText []byte
		key        []byte
	}
	key1, _ := hex.DecodeString(`889645f82002791df852847461de9824`)
	key2 := []byte(`1234567890123456`)
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "testCase-01",
			args: args{
				cipherText: []byte(`CqT/33f3jyoiYqT8MtxEFk3x2rlfhmgzhxpHqWosSj4d3hq2EbrtVyx2aLj565ZQNTcPrcDipnvpq/D/vQDaLKW70O83Q42zvR0//OfnYLcIjTPMnqa+SOhsjQrSdu66ySSORCAo`),
				key:        key1,
			},
			want: []byte(`{"ai":"test-accountId","name":"用户姓名","idNum":"371321199012310912"}`),
			wantErr: false,
		},
		{
			name: "testCase-02",
			args: args{
				cipherText: []byte(`DrtezM7/uXDSZb+IErPNdWCx/XcFbZt/ety4XASU/01rU4Lagzj0ZGLQd8sLVDxE6axtINauypc02XdZFC/aoXNiHvK5dIozXAlz3PRq6Yg+EnsHt+JIfnJWrnr177O75euMIelJncvbU34I1dkB8Y0O2gRg`),
				key:        key1,
			},
			want: []byte(`{"ai":"test-accountId","name":"用户姓名","idNum":"371321199012310912"}`),
			wantErr: false,
		},
		{
			name: "testCase-03",
			args: args{
				cipherText: []byte(`BAYHAAoJBQMCCwEImU09fpkzNl8JgrjUGZEG50qpxTfJIOU=`),
				key:        key2,
			},
			want: []byte(`haxqer1`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesGCMNoPaddingDecryptBase64Decode(tt.args.cipherText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesGCMNoPaddingDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesGCMNoPaddingDecrypt() got = %s, want %v", got, tt.want)
			}
		})
	}
}

func TestAesGCMNoPaddingEncryptBase64Encode(t *testing.T) {
	type args struct {
		origData []byte
		key      []byte
	}
	key1, _ := hex.DecodeString(`2836e95fcd10e04b0069bb1ee659955b`)
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "testCase-01",
			args: args{
				origData: []byte(`{"ai":"test-accountId","name":"用户姓名","idNum":"371321199012310912"}`),
				key:      key1,
			},
			want: nil,
			wantErr: false,
		},
		{
			name: "testCase-02",
			args: args{
				origData: []byte(`haxqer1`),
				key:      key1,
			},
			want: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesGCMNoPaddingEncryptBase64Encode(tt.args.origData, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesGCMNoPaddingEncryptBase64Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("%s", got)
		})
	}
}