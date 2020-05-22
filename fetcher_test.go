package gofunc

import (
	"fmt"
	"testing"
	"time"
)

func TestFetchWithTimeout(t *testing.T) {
	type args struct {
		timeout time.Duration
		url     string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "testCase001", args: args{
			timeout: 60 * time.Millisecond,
			url:     "http://www.baidu.com",
		}, wantErr: true},
		{name: "testCase002", args: args{
			timeout: 1000 * time.Millisecond,
			url:     "http://www.baidu.com",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchWithTimeout(tt.args.timeout, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchWithTimeout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				fmt.Printf("%+v", got)
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("FetchWithTimeout() got = %v, want %v", got, tt.want)
			//}
		})
	}
}