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
		{name: "testCase04", args: args{s: "3C:BD:3E:EF:49:95"}, want: "E9735B7866674151B2DA394839D6FBC2"},
		{name: "testCase05", args: args{s: "3CBD3EEF4995"}, want: "7D852B74D742C0D571FD670442D56322"},
		{name: "testCase06", args: args{s: "E037BF476833"}, want: "2F0428623E396CC222D5C537E9D1E67E"},
		{name: "testCase07", args: args{s: "E0B6559DAF51"}, want: "23E8EE4510F8E1D8907E8684127A350F"},
		{name: "testCase08", args: args{s: "E0:B6:55:9D:AF:51"}, want: "A99459806B674CD9D21B3D2043E63614"},
		{name: "testCase09", args: args{s: `{"gv":"1.0","sn":"AC37C2609FD24BDAC4D5639C280AE237","ra":"1","g":"20000010","v":"1.0","sid":"578","wf":"","no":"2","cu":"639"}sdb1!2@3#`}, want: "445A68F74623ADB8FB61F9CC7B31258A"},
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

