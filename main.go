package main

import (
	"github.com/kuochaoyi/gormpgwf/database"
	_ "github.com/kuochaoyi/gormpgwf/database"
	"log"
)

type Demo struct {
	database.BaseModel
	database.BaseModelSoftDelete
	database.BaseModelJsonb
	database.BaseModelSerialID
}

type ClassRoom struct {
	database.BaseModel
	database.BaseModelSoftDelete
	State string `gorm:"type:jsonb", sql:"type:JSONB::JSONB"`
	database.BaseModelSerialID
}

type j struct {
	age  int
	name string
}

func main() {
	// database.Open()
	// database.DBClient.Insert()

	db := database.DBClient.DB
	db.Debug().AutoMigrate(&database.BaseModel1{})

	user := &database.BaseModel{}
	if db.Create(&user).Error != nil {
		log.Panic("Unable to create user.")
	}

}
