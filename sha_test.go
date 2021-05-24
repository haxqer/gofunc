package gofunc

import "testing"

func TestSha1Upper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase01", args: args{s:"80-0A-80-5A-56-65"}, want: "A6F70E51D76296DF0C5F7B10B440F99D58D3CF60"},
		{name: "testCase01", args: args{s:"80-0A-80-5A-56-61"}, want: "ECCD8285D22947004353AE7287A6FED07D510638"},
		{name: "testCase01", args: args{s:""}, want: "DA39A3EE5E6B4B0D3255BFEF95601890AFD80709"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha1Upper(tt.args.s); got != tt.want {
				t.Errorf("Sha1Upper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSha1Upper(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Sha1Upper("http://www.mgtv.com/b/291325/2979284.html")
	}
}

func TestSha1Lower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase01", args: args{s:"80-0A-80-5A-56-65"}, want: "a6f70e51d76296df0c5f7b10b440f99d58d3cf60"},
		{name: "testCase02", args: args{s:"80-0A-80-5A-56-61"}, want: "eccd8285d22947004353ae7287a6fed07d510638"},
		{name: "testCase03", args: args{s:""}, want: "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha1Lower(tt.args.s); got != tt.want {
				t.Errorf("Sha1Upper() = %v, want %v", got, tt.want)
			}
		})
	}
}


func BenchmarkSha1Lower(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Sha1Lower("http://www.mgtv.com/b/291325/2979284.html")
	}
}

func TestSha256Lower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testCase01",
			args: args{s:`2836e95fcd10e04b0069bb1ee659955bappIdtest-appIdbizIdtest-bizIdidtest-idnametest-nametimestamps1584949895758{"data":"CqT/33f3jyoiYqT8MtxEFk3x2rlfhmgzhxpHqWosSj4d3hq2EbrtVyx2aLj565ZQNTcPrcDipnvpq/D/vQDaLKW70O83Q42zvR0//OfnYLcIjTPMnqa+SOhsjQrSdu66ySSORCAo"}`},
			want: "386c03b776a28c06b8032a958fbd89337424ef45c62d0422706cca633d8ad5fd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha256Lower(tt.args.s); got != tt.want {
				t.Errorf("Sha256Lower() = %v, want %v", got, tt.want)
			}
		})
	}
}