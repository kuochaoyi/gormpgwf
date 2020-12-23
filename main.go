package main

import (
	"goland_gorm/database"
	_ "goland_gorm/database"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// database.Open()
	// database.DBClient.Insert()
	db := database.DBClient.DB
	json := `{"age":  27, "name": "Yan"}`

	db.AutoMigrate(&database.BaseModelSoftDelete{})
	db.Create(&database.BaseModelSoftDelete{
		BaseModel: database.BaseModel{
			ID:        uuid.UUID{},
			CreatedAt: time.Time{},
			UpdatedAt: nil,
			BaseModelJsonb: database.BaseModelJsonb{
				JsonStore: postgres.Jsonb{
					RawMessage: []byte(json),
				},
			},
		},
		DeletedAt: nil,
	})

}
