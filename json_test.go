package go_func

import (
	"reflect"
	"testing"
)

func TestMarshalFalseEscapeHTML(t *testing.T) {
	type args struct {
		t interface{}
	}
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"
	countryCapitalMap["and"] = "&"

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "testCase1",
			args:    args{countryCapitalMap},
			want:    []byte(`{"France":"巴黎","India":"新德里","Italy":"罗马","Japan":"东京","and":"&"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalFalseEscapeHTML(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalFalseEscapeHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalFalseEscapeHTML() got = %s, want %v", got, tt.want)
			}
		})
	}
}
