package gofunc

import (
	"testing"
	"time"
)

func TestIsCommonWeekend(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name:"testCase-01", args: args{t: time.Date(2021, 5, 30, 0, 0, 0, 0, time.UTC)}, want: true},
		{name:"testCase-02", args: args{t: time.Date(2021, 5, 31, 0, 0, 0, 0, time.UTC)}, want: false},
		{name:"testCase-03", args: args{t: time.Date(2021, 5, 29, 0, 0, 0, 0, time.UTC)}, want: true},
		{name:"testCase-04", args: args{t: time.Date(2021, 5, 28, 0, 0, 0, 0, time.UTC)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCommonWeekend(tt.args.t); got != tt.want {
				t.Errorf("IsCommonWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChineseHoliday(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name:"testCase-01", args: args{t: time.Date(2021, 5, 30, 0, 0, 0, 0, time.UTC)}, want: true},
		{name:"testCase-02", args: args{t: time.Date(2021, 5, 31, 0, 0, 0, 0, time.UTC)}, want: false},
		{name:"testCase-03", args: args{t: time.Date(2021, 5, 29, 0, 0, 0, 0, time.UTC)}, want: true},
		{name:"testCase-04", args: args{t: time.Date(2021, 5, 28, 0, 0, 0, 0, time.UTC)}, want: false},
		{name:"testCase-05", args: args{t: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)}, want: true},
		{name:"testCase-06", args: args{t: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)}, want: false},
		{name:"testCase-07", args: args{t: time.Date(2021, 4, 25, 0, 0, 0, 0, time.UTC)}, want: false},
		{name:"testCase-08", args: args{t: time.Date(2021, 8, 18, 0, 0, 0, 0, time.UTC)}, want: false},
		{name:"testCase-09", args: args{t: time.Date(2021, 9, 18, 0, 0, 0, 0, time.UTC)}, want: false},
		{name:"testCase-10", args: args{t: time.Date(2021, 9, 19, 0, 0, 0, 0, time.UTC)}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChineseHoliday(tt.args.t); got != tt.want {
				t.Errorf("IsChineseHoliday() = %v, want %v", got, tt.want)
			}
		})
	}
}