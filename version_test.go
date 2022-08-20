package gofunc

import "testing"

// @author: Haxqer
// @email: haxqer666@gmail.com
// @since: 2022/8/21
// @desc: TODO

func TestVersionToInt64(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{name: "testCase-01", args: args{version: "1.2.3.4"}, want: 1000200030004, wantErr: false},
		{name: "testCase-02", args: args{version: "1.2.3.4.5"}, want: 0, wantErr: true},
		{name: "testCase-03", args: args{version: "1"}, want: 1, wantErr: false},
		{name: "testCase-04", args: args{version: "9999"}, want: 9999, wantErr: false},
		{name: "testCase-05", args: args{version: "9999.888"}, want: 99990888, wantErr: false},
		{name: "testCase-06", args: args{version: "107.0.5251.1"}, want: 107000052510001, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VersionToInt64(tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("VersionToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VersionToInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkVersionToInt64(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		VersionToInt64("107.0.5251.1")
	}
}

func TestVersionInt64ToStr(t *testing.T) {
	type args struct {
		versionInt int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "testCase-01", args: args{versionInt: 1000200030004}, want: "1.2.3.4", wantErr: false},
		{name: "testCase-02", args: args{versionInt: 0}, want: "", wantErr: true},
		{name: "testCase-03", args: args{versionInt: 1}, want: "1", wantErr: false},
		{name: "testCase-04", args: args{versionInt: 100101}, want: "10.101", wantErr: false},
		{name: "testCase-04", args: args{versionInt: 107000052510001}, want: "107.0.5251.1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VersionInt64ToStr(tt.args.versionInt)
			if (err != nil) != tt.wantErr {
				t.Errorf("VersionInt64ToStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VersionInt64ToStr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkVersionInt64ToStr(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		VersionInt64ToStr(107000052510001)
	}
}
