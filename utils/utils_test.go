package utils

import (
	"goland_gorm/database"
	_ "goland_gorm/database"
	"testing"
)

func TestGetTodaySerial(t *testing.T) {
	db := database.DBClient.DB
	db.AutoMigrate()

	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "now",
			want: "20201229",
		},
		{
			name: "2020-12-25",
			want: "20201225",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTodaySerial(); got != tt.want {
				t.Errorf("GetTodaySerial() = %v, want %v", got, tt.want)
			}
		})
	}
}
