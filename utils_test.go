package gofunc

import (
	"reflect"
	"testing"
)

func TestPack(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "testCase01", args: args{
			b: []byte("sadfsdfasdfds"),
		},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pack(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			b, err := Unpack(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(b, tt.args.b) {
				t.Errorf("Pack() got = %v, want %v", b, tt.args.b)
			}
		})
	}
}