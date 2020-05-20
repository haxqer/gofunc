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