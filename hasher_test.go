package gofunc

import "testing"

func TestMd5Upper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase01", args: args{s: "123"}, want: "202CB962AC59075B964B07152D234B70"},
		{name: "testCase02", args: args{s: "http://www.mgtv.com/b/291325/2979284.html"}, want: "D0E2109A8FF867EE55F83F8765273FF9"},
		{name: "testCase03", args: args{s: "abc"}, want: "900150983CD24FB0D6963F7D28E17F72"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Upper(tt.args.s); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMd5Upper(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Md5Upper("http://www.mgtv.com/b/291325/2979284.html")
	}
}

func TestMd5Lower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase01", args: args{s: "123"}, want: "202cb962ac59075b964b07152d234b70"},
		{name: "testCase02", args: args{s: "http://www.mgtv.com/b/291325/2979284.html"}, want: "d0e2109a8ff867ee55f83f8765273ff9"},
		{name: "testCase03", args: args{s: "abc"}, want: "900150983cd24fb0d6963f7d28e17f72"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Lower(tt.args.s); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMd5Lower(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Md5Lower("http://www.mgtv.com/b/291325/2979284.html")
	}
}

