package xxtea

import "testing"

func TestDecryptString(t *testing.T) {
	type args struct {
		str string
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "testCase-01",
			args: args{
				str: "77BY4stkmnKD8UqMJHW4AwJVkSraIs6JzpWsxMw1It1DrrtxuNxCwJ/uqrRP6VhAiALl4FFTfcM9hk/JPA69+G7CX5bXBhJRo623JitDtwI6q5xNu2DXoDY1eY6elJY0T+r2KEFHWWnPuGSBcA/fstYR2zz0bv8VvI3ka9ahy+wVdN4juO7n8sNaow5gaXAi8DI0ME6r2ss1QDuGYbbcg/VbLWLM3P3srnNvMFIrVcSVC8kOUqhPD1+bi11KqljD2aJw9KpmKyvURQf0JnSWOfLidX3gEbGnCraummg9cSUuvAv2i6HX+l47gGyhIlA1IFFj/PyY22fCS2wSyQZvN/93EoWR1T9iC2S1iHwwKkBY5lNQZksGc4ZrwxPrXyoxROODQfaSr+/Ih20e9wJjALGBsI8tMxGJOt8LiCZf8He2lNYeIcri8S8WS42bxBjyHtSUShAKW41JTO7+a2oDXnJ/Nz7nC2Npyt07SuEDBOQ+eFzoxvy14bP9ozehZMAf25eF1QgA8oxzEl9CV6KCPJNQANlPMgmlNa113CmZftR/54Fc1zZxk556qV0ozm4aw+S4yZTCID53KXf3cRJZ4fq9TygnSshMxiqWfE/GE2bLYMenajcusI5UMVIQXEv5tzah1w==",
				key: "sdb123",
			},
			want: `{"code":"200","msg":"成功","responseTm":"20210519145844","data":{"userName":"yy292084","game_name":"道天录","litle_url":null,"privacy_state":"aHR0cDovLj3d3dy5wa2V5LmNuL3BwL3ByaXZhY3kuaHRtbA==","register_state":"aHR0cDovLo3d3dy5wa2V5LmNuL3BwL3Byb3RvY29sLmh0bWw=","sdk_account":null,"sdk_passwd":null,"game_version_verify":"0","generalize_switch":0,"code":null,"msg":null,"responseTm":"20210519145844","csgroup":"856104850","csqq":"519363413","succeed":false},"succeed":true}`,
			wantErr: false,
		},
		{
			name: "testCase-02",
			args: args{
				str: "PQd4b1V1mvMra6/4qpviMdqinaRCkkWwHMW8p43Ham5tarUfCEFvTb7x3R03+UKbXFSk9A==",
				key: "sdb123",
			},
			want: `{"code":400,"msg":"请求参数错误","data":4}`,
			wantErr: false,
		},
		{
			name: "testCase-03",
			args: args{
				str: "/San8SE+R5CB4EedL9i2BiEr7VSx2nIHIYkqYeVx4eA85dw9psPwRn4cPBr1naZVWbsoiK1ZfE9u54qyYAJoJo5XWfk=",
				key: "sdb123",
			},
			want: `{"code":400,"msg":"请求参数错误","data":"加解密错误"}`,
			wantErr: false,
		},
		{
			name: "testCase-04",
			args: args{
				str: "5uANLaE7PEGMZigg3BJ+Ts983og7WiRURN5mU9MPQEn++9o7CEGpEOw/sfztXS/zJHW2yNQvHQ/yN/MX",
				key: "1234567890adfsjd",
			},
			want: `{"code":400,"msg":"请求参数错误","data":"加解密错误"}`,
			wantErr: false,
		},
		{
			name: "testCase-05",
			args: args{
				str: "XXfPte0LqbUDr/FcYSCxhSPpF6SocPH1NZqAGm9CTeXVjGgqgVYdb9EfD0CEhfvQeg6yXhvFQNxuSqdFyqNYX9rMr+S85O3nppokKQ==",
				key: "1234567890adfsj1",
			},
			want: `{"code":400,"msg":"请求参数错误","data":"加解密错误"}`,
			wantErr: false,
		},
		{
			name: "testCase-06",
			args: args{
				str: "xETOidnksnb6ch6k00qvAj1SAoHe2Q6Qu9VKTPpb4Ioi0Old5b9R0pDIAZqrC+TBxRbC00qvo06N/cNdkqwL2rogAhRQX80KcsVIbng0W6+4D8zrw9aT4dy86yBfE5Afb7ygQxA0aNxYWC4WgyzRPQ==",
				key: "1234567890adfsjd",
			},
			want: ``,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecryptString(tt.args.str, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecryptString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptString(t *testing.T) {
	type args struct {
		str string
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testCase-01",
			args: args{
				str: `{"auth_platform": "APPLE","apple_user_id": "001876.345341e39ae64c6a92ee87dbbb2ab1c4.0248","apple_identity_token": "eyJraWQiOiJZdXlYb1kiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLnlvdWppdS5xaWNhaW1hamlhbmciLCJleHAiOjE2MjE2NTE3MzEsImlhdCI6MTYyMTU2NTMzMSwic3ViIjoiMDAxODc2LjM0NTM0MWUzOWFlNjRjNmE5MmVlODdkYmJiMmFiMWM0LjAyNDgiLCJjX2hhc2giOiJhVU5lTm01T2JJeXI2bnRwM3IwdWlnIiwiZW1haWwiOiI3OTMzMzI4MTZAcXEuY29tIiwiZW1haWxfdmVyaWZpZWQiOiJ0cnVlIiwiYXV0aF90aW1lIjoxNjIxNTY1MzMxLCJ0cmFuc2Zlcl9zdWIiOiIwMDE4NzYucmRjOTU1ZTNkY2E3ZjRjNjU4ZmJhYmIwMWE4ZDNhMjFhIiwibm9uY2Vfc3VwcG9ydGVkIjp0cnVlfQ.cjCY_YmhLGf0wL8pTZDlx01oarj2SV8eUUrs1Moww8kfp0eWmwbFRC5s4hEHBfYFecy4lELW1bhfe9YjGM0GkwO96n4AbhSSxvbo8mQLFWCC-YRRNZZ2UlboIpZR1jTBlDYm9Yx-idUWCVTBgFacaaEWWHqV3Tgb3zjPfB6uMvq2xTU3vzD3HUYwmpB8m-1uoTPkROHK4gvutIZg85SVr1ZQ9WL7yEBt0InwzVrtVTx44BE8O2W-dJA7CAb6GJzsLIc3uipjnp1K7w2gAos_gUl5F0HGJafwyGnuYfYsAmZ8u4WK4LQsTyD_smONUiHYF5AAYPVPSaq4iUzJt98DuQ"}`,
				key: "sdb123",
			},
			want: ``,
		},
		{
			name: "testCase-02",
			args: args{
				str: `{"game_version": "123"}`,
				key: "sdb123",
			},
			want: ``,
		},
		{
			name: "testCase-03",
			args: args{
				str: `{"pay_method":"WECHAT_H5","order_amount":2,"game_product_name":"测试游戏商品名称","game_product_amount":1}`,
				key: "sdb123",
			},
			want: `D1CO2WljB2nzbv7aUiGB1vQ2hGKN05P8aDAsYYFb8IqXPyc8fHGcKfC+YNXMw5mN9xsGZpVmzoc+7yWlGZ/OfMJKYfZdJItrbUjXWyLkBMgiu7DZV2l3S+9mFHjv0/HTOa3y9gYf29gv/5a8BI1chVjNVbhujTpy`,
		},
		{
			name: "testCase-04",
			args: args{
				str: `{"gss_app_id":"6002"}`,
				key: "51f85ee0687a3f21",
			},
			want: `XBXfrxu+rcKs7Bd4MRSNAsmAgYGzNPkKn4mUKLwY+bPhZX+WdWbqIg==`,
		},
		{
			name: "testCase-05",
			args: args{
				str: `{"auth_platform":"APPLE","apple_user_id":"001876.345341e39ae64c6a92ee87dbbb2ab1c4.0248","apple_identity_token":"eyJraWQiOiJZdXlYb1kiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLnlvdWppdS5xaWNhaW1hamlhbmciLCJleHAiOjE2MjE2NTE3MzEsImlhdCI6MTYyMTU2NTMzMSwic3ViIjoiMDAxODc2LjM0NTM0MWUzOWFlNjRjNmE5MmVlODdkYmJiMmFiMWM0LjAyNDgiLCJjX2hhc2giOiJhVU5lTm01T2JJeXI2bnRwM3IwdWlnIiwiZW1haWwiOiI3OTMzMzI4MTZAcXEuY29tIiwiZW1haWxfdmVyaWZpZWQiOiJ0cnVlIiwiYXV0aF90aW1lIjoxNjIxNTY1MzMxLCJ0cmFuc2Zlcl9zdWIiOiIwMDE4NzYucmRjOTU1ZTNkY2E3ZjRjNjU4ZmJhYmIwMWE4ZDNhMjFhIiwibm9uY2Vfc3VwcG9ydGVkIjp0cnVlfQ.cjCY_YmhLGf0wL8pTZDlx01oarj2SV8eUUrs1Moww8kfp0eWmwbFRC5s4hEHBfYFecy4lELW1bhfe9YjGM0GkwO96n4AbhSSxvbo8mQLFWCC-YRRNZZ2UlboIpZR1jTBlDYm9Yx-idUWCVTBgFacaaEWWHqV3Tgb3zjPfB6uMvq2xTU3vzD3HUYwmpB8m-1uoTPkROHK4gvutIZg85SVr1ZQ9WL7yEBt0InwzVrtVTx44BE8O2W-dJA7CAb6GJzsLIc3uipjnp1K7w2gAos_gUl5F0HGJafwyGnuYfYsAmZ8u4WK4LQsTyD_smONUiHYF5AAYPVPSaq4iUzJt98DuQ","ts":1111}`,
				key: "1234567890adfsjd",
			},
			want: `DmuggcRtB4y1z2YKstJwdpo4NS/ro+yDwFhfQBsNK85KLY8v3nUgGT0+t1espcqTM76skF/G5y9gyomcj7iGpx4od0JLCWIOBjhLtQJoDsiFRFPabLWs0isSCM4VoAmN/iwjdIe6ATOjlj2Q0XZVla1Z/btoTSRs8scqLQ==`,
		},
		{
			name: "testCase-05",
			args: args{
				str: `{"pay_method":"APPLE_NATIVE","order_amount":"1","game_product_name":"test","game_product_amount":"1"}`,
				key: "1234567890adfsjd",
			},
			want: `DmuggcRtB4y1z2YKstJwdpo4NS/ro+yDwFhfQBsNK85KLY8v3nUgGT0+t1espcqTM76skF/G5y9gyomcj7iGpx4od0JLCWIOBjhLtQJoDsiFRFPabLWs0isSCM4VoAmN/iwjdIe6ATOjlj2Q0XZVla1Z/btoTSRs8scqLQ==`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptString(tt.args.str, tt.args.key); got != tt.want {
				t.Errorf("EncryptString() = %v, want %v", got, tt.want)
			}
		})
	}
}