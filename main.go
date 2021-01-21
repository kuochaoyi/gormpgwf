package main

import (
	"github.com/kuochaoyi/gormpgwf/database"
	_ "github.com/kuochaoyi/gormpgwf/database"
	"log"
)

func main() {
	// database.Open()
	// database.DBClient.Insert()

	db := database.DBClient.DB
	db.Debug().AutoMigrate(&database.BaseModel{})

	user := &database.BaseModel{}
	if db.Create(&user).Error != nil {
		log.Panic("Unable to create user.")
	}
}
