package utils

import "testing"

func TestGetTodaySerial(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTodaySerial(); got != tt.want {
				t.Errorf("GetTodaySerial() = %v, want %v", got, tt.want)
			}
		})
	}
}
