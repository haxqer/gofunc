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
		{name: "testCase10", args: args{s: `{"auth_platform": "WECHAT","oauth_code": "AAAAAAAAAAAAAAAAAAAAAAAA"}sdb1!2@3#`}, want: "DCFE51932C2BEF233E352CF22BADF6FF"},
		{name: "testCase11", args: args{s: `{"auth_platform": "APPLE","apple_user_id": "001876.345341e39ae64c6a92ee87dbbb2ab1c4.0248","apple_identity_token": "eyJraWQiOiJZdXlYb1kiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLnlvdWppdS5xaWNhaW1hamlhbmciLCJleHAiOjE2MjE2NTE3MzEsImlhdCI6MTYyMTU2NTMzMSwic3ViIjoiMDAxODc2LjM0NTM0MWUzOWFlNjRjNmE5MmVlODdkYmJiMmFiMWM0LjAyNDgiLCJjX2hhc2giOiJhVU5lTm01T2JJeXI2bnRwM3IwdWlnIiwiZW1haWwiOiI3OTMzMzI4MTZAcXEuY29tIiwiZW1haWxfdmVyaWZpZWQiOiJ0cnVlIiwiYXV0aF90aW1lIjoxNjIxNTY1MzMxLCJ0cmFuc2Zlcl9zdWIiOiIwMDE4NzYucmRjOTU1ZTNkY2E3ZjRjNjU4ZmJhYmIwMWE4ZDNhMjFhIiwibm9uY2Vfc3VwcG9ydGVkIjp0cnVlfQ.cjCY_YmhLGf0wL8pTZDlx01oarj2SV8eUUrs1Moww8kfp0eWmwbFRC5s4hEHBfYFecy4lELW1bhfe9YjGM0GkwO96n4AbhSSxvbo8mQLFWCC-YRRNZZ2UlboIpZR1jTBlDYm9Yx-idUWCVTBgFacaaEWWHqV3Tgb3zjPfB6uMvq2xTU3vzD3HUYwmpB8m-1uoTPkROHK4gvutIZg85SVr1ZQ9WL7yEBt0InwzVrtVTx44BE8O2W-dJA7CAb6GJzsLIc3uipjnp1K7w2gAos_gUl5F0HGJafwyGnuYfYsAmZ8u4WK4LQsTyD_smONUiHYF5AAYPVPSaq4iUzJt98DuQ"}sdb1!2@3#`}, want: "F0CA9394E01DA824B8437D3C9E9350A2"},
		{name: "testCase12", args: args{s: `{"game_version": "123"}sdb1!2@3#`}, want: "B6B5BA2AA9A58488E2425FBAFD37F488"},
		{name: "testCase13", args: args{s: `{"pay_method":"WECHAT_H5","order_amount":2,"game_product_name":"测试游戏商品名称","game_product_amount":1}sdb1!2@3#`}, want: "3C3F604C9ACDD686FA860ED5CF317DE3"},
		{name: "testCase14", args: args{s: `{"gss_app_id":"6002"}sdb1!2@3#`}, want: "58699F25CC4B87939305FA53F52F154D"},
		{name: "testCase14", args: args{s: `51f85ee0687a3f21{"a":1}`}, want: "2DE4C13EDB58035857A9B153BB24A3F9"},
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

