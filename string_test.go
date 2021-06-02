package gofunc

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	code := m.Run()
	os.Exit(code)
}

func TestRandStringRunes(t *testing.T) {
	type args struct {
		n      int
		strArr []rune
	}
	tests := []struct {
		name      string
		args      args
		wantlenth int
	}{
		{name: "testCase01", args: args{
			n:      11,
			strArr: []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"),
		}, wantlenth: 11},
		{name: "testCase02", args: args{
			n:      11,
			strArr: []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"),
		}, wantlenth: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandStringRunes(tt.args.n, tt.args.strArr)
			if len(got) != tt.wantlenth {
				t.Errorf("len(RandStringRunes()) = %d, want %v", len(got), tt.wantlenth)
			}
			fmt.Printf("%s\n", got)
		})
	}
}

func BenchmarkRandStringRunes(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RandStringRunes(11, []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"))
	}
}

func TestBase64Encode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase01", args: args{s: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"},
			want: "QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMjM0NTY3ODk=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Encode(tt.args.s); got != tt.want {
				t.Errorf("Base64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Base64Encode("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")
	}
}

func TestBase64Decode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "testCase01", args: args{s: "QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMjM0NTY3ODk="},
			want:    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64Decode(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base64Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBase64Decode(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := Base64Decode("QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMjM0NTY3ODk=")
		if err != nil {
			b.Errorf("Base64Decode() error = %v ", err)
		}
	}
}

func TestBase64DecodeByte(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "testCase01", args: args{b: []byte("QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMjM0NTY3ODk=")},
			want:    []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64DecodeByte(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64DecodeByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base64DecodeByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRFCEncode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase01", args: args{s: "sfadf fadsdfa+Sfd"}, want: "sfadf+fadsdfa%2BSfd"},
		{name: "testCase02", args: args{s: ` "#%&()+,/:;<=>?@\|{}`}, want: "+%22%23%25%26%28%29%2B%2C%2F%3A%3B%3C%3D%3E%3F%40%5C%7C%7B%7D"},
		{name: "testCase03", args: args{s: `Dalvik/2.1.0 (Linux; U; Android 9; Nokia X6 Build/PPR1.180610.011)`},
			want: "Dalvik%2F2.1.0+%28Linux%3B+U%3B+Android+9%3B+Nokia+X6+Build%2FPPR1.180610.011%29"},
		{name: "testCase04", args: args{s: "sfadf 这是中文"}, want: "sfadf+%E8%BF%99%E6%98%AF%E4%B8%AD%E6%96%87"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RFCEncode(tt.args.s); got != tt.want {
				t.Errorf("EncodeParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRFCEncode(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RFCEncode(` "#%&()+,/:;<=>?@\|{}`)
	}
}

func TestRFCCheck(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "testCase01", args: args{s: "sfadf fadsdfa+Sfd"}, want: false},
		{name: "testCase02", args: args{s: "sfadffadsdfa+Sfd"}, want: true},
		{name: "testCase03", args: args{s: "sfadffadsdfa+Sfd中文"}, want: false},
		{name: "testCase04", args: args{s: `Dalvik/2.1.0 (Linux; U; Android 9; Nokia X6 Build/PPR1.180610.011)`}, want: false},
		{name: "testCase05", args: args{s: `Dalvik%2F2.1.0+%28Linux%3B+U%3B+Android+9%3B+Nokia+X6+Build%2FPPR1.180610.011%29`}, want: true},
		{name: "testCase05", args: args{s: ` "#%&()+,/:;<=>?@\|{}`}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RFCCheck(tt.args.s); got != tt.want {
				t.Errorf("RFCCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRFCCheck(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RFCCheck(`Dalvik%2F2.1.0+%28Linux%3B+U%3B+Android+9%3B+Nokia+X6+Build%2FPPR1.180610.011%29`)
	}
}

func TestBase64EncodeByte(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "testCase01", args: args{b: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")},
			want:    []byte("QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMjM0NTY3ODk="),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64EncodeByte(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base64EncodeByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexEncode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase01", args: args{s: "asdfasfsf13"}, want: "6173646661736673663133"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexEncode(tt.args.s); got != tt.want {
				t.Errorf("HexEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "testCase01", args: args{s: "6173646661736673663133"}, want: []byte("asdfasfsf13"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexDecode(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexDecode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexEncodeByte(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "testCase01", args: args{b: []byte("asdfasfsf13")}, want: []byte("6173646661736673663133")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexEncodeByte(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexEncodeByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexDecodeByte(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "testCase01", args: args{b: []byte("6173646661736673663133")}, want: []byte("asdfasfsf13"), wantErr: false},
		{name: "testCase01", args: args{b: []byte("asdfzxcvdf6173646661736673663133")}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexDecodeByte(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexDecodeByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexDecodeByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToString(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase-01", args: args{i: 1234567890123111}, want: "1234567890123111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ToString(tt.args.i); got != tt.want {
				t.Errorf("Int64ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestStringToInt64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{name: "testCase-01", args: args{s: "99999999999"}, want: 99999999999, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase-01", args: args{i: 999999}, want: "999999"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToString(tt.args.i); got != tt.want {
				t.Errorf("IntToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "testCase-01", args: args{s: "99999999"}, want: 99999999, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}