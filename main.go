package main

import (
	"log"

	"github.com/kuochaoyi/gormpgwf/database"
	_ "github.com/kuochaoyi/gormpgwf/database"
)

func main() {
	// database.Open()
	// database.DBClient.Insert()

	db := database.DBClient.DB
	// db.Debug().AutoMigrate(&database.BaseModel{})

	// user := &database.BaseModel{}
	// if db.Create(&user).Error != nil {
	// 	log.Panic("Unable to create user.")
	// }

	var pNo string
	db.Raw("SELECT p_order_no_new()").Scan(&pNo)
	log.Printf("no: %s", pNo)

	order := `{
		{bcd16bb9-1f72-47af-90af-2e8ce2cf0668, 15, 300},
		{444d77a1-cb11-4444-a82c-dd71645f4858, 30, 100}
		}`
	var created int
	db.Raw("SELECT p_insert(?)", order).Scan(&created)
	log.Printf("created: %v, %T", created, created)
}
