package gofunc

import (
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"testing"
)


var privateKeyBytes = []byte(`-----BEGIN RSA PRIVATE KEY-----
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

var publicKeyBytes = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCyBUxPA9X3tdJSsgNi0/CctBsB
bvJEUgt5FWbbG4nTp2r7XK2TVk/YBZ6jq7kiA1rzJ/xmJTdpUdVUhPfX1DN/7iap
b5K3z/NhIO/jsJgGO+YtgK5IWEcjGwElPZnOsMk6iNZAWCGa7EyA0FHkul7w4eFO
jC+RGqbKfsl306EUtwIDAQAB
-----END PUBLIC KEY-----
`)

func Test_RsaDecode(t *testing.T) {
	type args struct {
		input []byte
		key   []byte
	}
	key, _ := hex.DecodeString("58F8925C485ECF941241CEF4C25819AA1CC67F52ED2C62CEDFAECAF72B420A445548E4D8BC246A114FEB7BB2EE520F32DD65DEBFAD13077BB09B08BDD9428F24B112306ECFB612579A2305E8212109F6833889F9B0552B3D5CD11C50B35B69DF9691E9AE3264930C29774C358A271853ADAC2C138D7F371F080E9E86E0A8CAE6")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name:"testCase01", args:args{
			input: key,
			key:   privateKeyBytes,
		}, want:nil, wantErr:false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RsaDecode(tt.args.input, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RsaDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("%X \n", got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("RsaDecode() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func BenchmarkRsaDecode(b *testing.B) {
	key, _ := hex.DecodeString("58F8925C485ECF941241CEF4C25819AA1CC67F52ED2C62CEDFAECAF72B420A445548E4D8BC246A114FEB7BB2EE520F32DD65DEBFAD13077BB09B08BDD9428F24B112306ECFB612579A2305E8212109F6833889F9B0552B3D5CD11C50B35B69DF9691E9AE3264930C29774C358A271853ADAC2C138D7F371F080E9E86E0A8CAE6")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := RsaDecode(key, privateKeyBytes)
		if err != nil {
			b.Errorf("RsaDecode() err %+v = got = %s", err, got)
		}
	}
}

func BenchmarkRsaDecodeWithPrivate(b *testing.B) {
	key, _ := hex.DecodeString("58F8925C485ECF941241CEF4C25819AA1CC67F52ED2C62CEDFAECAF72B420A445548E4D8BC246A114FEB7BB2EE520F32DD65DEBFAD13077BB09B08BDD9428F24B112306ECFB612579A2305E8212109F6833889F9B0552B3D5CD11C50B35B69DF9691E9AE3264930C29774C358A271853ADAC2C138D7F371F080E9E86E0A8CAE6")
	rsaKey, err := ParsePrivateKeyFromKeyBytes(privateKeyBytes)
	if err != nil {
		b.Errorf("ParsePrivateKeyFromKeyBytes() err %+v ", err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := RsaDecodeWithPrivate(key, rsaKey)
		if err != nil {
			b.Errorf("RsaDecodeWithPrivate() err %+v = got = %s", err, got)
		}
	}
}

func BenchmarkRsaEncodeWithPublic(b *testing.B) {
	ekey, _ := hex.DecodeString("DDF11F26FC4321EB0C2B83C9D1B2120B64716E4FF28FB3D0E8F4FF73B381A4CC")
	rsaKey, err := ParsePublicKeyFromKeyBytes(publicKeyBytes)
	if err != nil {
		b.Errorf("ParsePublicKeyFromKeyBytes() err %+v ", err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, err := RsaEncodeWithPublic(ekey, rsaKey)
		if err != nil {
			b.Errorf("RsaEncodeWithPublic() err %+v = got = %s", err, got)
		}
	}
}

func TestRsaEncodeWithPublic(t *testing.T) {
	type args struct {
		input  []byte
		rsaKey *rsa.PublicKey
	}
	ekey, _ := hex.DecodeString("DDF11F26FC4321EB0C2B83C9D1B2120B64716E4FF28FB3D0E8F4FF73B381A4CC")
	rsaKey, err := ParsePublicKeyFromKeyBytes(publicKeyBytes)
	if err != nil {
		t.Errorf("ParsePublicKeyFromKeyBytes() err %+v ", err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name:"testCase01", args:args{
			input: ekey,
			rsaKey:   rsaKey,
		}, want:nil, wantErr:false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RsaEncodeWithPublic(tt.args.input, tt.args.rsaKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("RsaEncodeWithPublic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("RsaEncodeWithPublic() got = %v, want %v", got, tt.want)
			//}
			fmt.Printf("%X \n", got)
		})
	}
}