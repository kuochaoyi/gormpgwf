package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kuochaoyi/gormpgwf/database"

	_ "github.com/kuochaoyi/gormpgwf/database"
)

const (
	serial_add    = 1
	layout_data   = "20060102"
	LAYOUT_DIGITS = "000000"
	// layout_digit  = "6"
)

type SerialGenerator interface {
	SetSerial(table string) string // Set
	Serial(table string) string    // Get newly serial
}

type SerialPgHandler struct {
}

//
func (h *SerialPgHandler) SetSerial(table string) string {
	latestID := h.Serial(table)

	i, _ := strconv.Atoi(latestID[8:14])
	return MakeSerial(i + serial_add)
}

// Get SerialID == today max
func (h *SerialPgHandler) Serial(table string) string {
	var latestID string
	db := database.DBClient.DB

	/*
		db.Raw("SELECT MAX(serial_id) FROM ? WHERE created_at = CURRENT_DATE ", tableName).Scan(&latestID)
		ERROR: syntax error at or near "$1" (SQLSTATE 42601); ERROR: syntax error at or near "$1" (SQLSTATE 42601)
		[0.565ms] [rows:-] SELECT MAX(serial_id) FROM 'base_models' WHERE created_at = CURRENT_DATE
	*/

	/*
		select serial_id from base_models where
			 created_at = (
		select max(created_at) from base_models
			 );
	*/
	whereqy := "created_at = ( \nselect max(created_at) from " + table + "\n\t )"
	row := db.Table(table).Where(whereqy).Select("serial_id").Row()
	row.Scan(&latestID)
	if latestID != "" {
		return latestID
	}

	return GetTodaySerial() + LAYOUT_DIGITS
}

// Get today to serial. ex.20200123
func GetTodaySerial() string {
	return time.Now().Format(layout_data)
}

// Default: 6 digit => 000001
func FormatSerial(s int) string {
	return fmt.Sprintf("%06d", s)
}

func MakeSerial(i int) string {
	return GetTodaySerial() + FormatSerial(i)
}
