package gofunc

import "testing"

func TestSha1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase01", args: args{s:"80-0A-80-5A-56-65"}, want: "A6F70E51D76296DF0C5F7B10B440F99D58D3CF60"},
		{name: "testCase01", args: args{s:"80-0A-80-5A-56-65"}, want: "A6F70E51D76296DF0C5F7B10B440F99D58D3CF60"},
		{name: "testCase01", args: args{s:""}, want: "DA39A3EE5E6B4B0D3255BFEF95601890AFD80709"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha1(tt.args.s); got != tt.want {
				t.Errorf("Sha1() = %v, want %v", got, tt.want)
			}
		})
	}
}