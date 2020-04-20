package gofunc

import (
	"fmt"
	"math/rand"
	"os"
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
		name string
		args args
		wantlenth int
	}{
		{name: "testCase01", args: args{
			n: 11,
			strArr: []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"),
		}, wantlenth: 11},
		{name: "testCase02", args: args{
			n: 11,
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

