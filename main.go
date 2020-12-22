package main

import (
	"goland_gorm/database"
	_ "goland_gorm/database"
)

func main() {
	// database.Open()
	// database.DBClient.Insert()
	database.DBClient.DB.AutoMigrate(&database.BaseModelSoftDelete{})
	database.DBClient.DB.Create(&database.BaseModelSoftDelete{})
}
