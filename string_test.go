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

func TestBase64Decode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "testCase01", args: args{s: "QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMjM0NTY3ODk="},
			want:    []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
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
