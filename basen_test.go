package gofunc

import "testing"

func TestDecimalToAny(t *testing.T) {
	type args struct {
		in     int
		strArr []rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase01", args: args{
			in: 1586966404 / 4,
			strArr: []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"),
		}, want: "Xpcvh"},
		{name: "testCase02", args: args{
			in: 1556467130 / 4,
			strArr: []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"),
		}, want: "XMXNu"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalToAny(tt.args.in, tt.args.strArr); got != tt.want {
				t.Errorf("DecimalToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDecimalToAny(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DecimalToAny(1586966404 / 4, []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"))
	}
}

