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
		},
		{
			name: "testCase-01",
			args: args{
				str: "iJrP2+hLwVtfMgsxaRRhuAPTnS1ilzt4VNM7iT54LXYQX3pgIOog+FCEcrhugKwff90lGI2ZkUmJpOw8ONVwxRN2xnpc6uadlfLLX192sSm/YLo6sTkx/ZBX8IWzmXEbo+XQbgdTS4Gubga0n7ki3d5p1yaoQ7DV50W3/dbiN6Pj6JY2",
				key: "sdb123",
			},
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
