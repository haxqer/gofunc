package gofunc

import "testing"

func TestTs(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := Ts(); got != tt.want {
				t.Errorf("Ts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEndOfTime(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{name: "testCase-01", want: 111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEndOfTime(); got != tt.want {
				t.Errorf("GetEndOfTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
