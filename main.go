package main

import (
	"encoding/json"
	"goland_gorm/database"
	_ "goland_gorm/database"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// database.Open()
	// database.DBClient.Insert()

	db := database.DBClient.DB
	db.AutoMigrate(&database.BaseModel{})

	// json := `{"age":  27, "name": "Yan"}`
	json1 := `{
	"name": "Pasta",
		"ingredients": ["Flour", "Eggs", "Salt", "Water"],
"organic": true,
"dimensions": {
"weight": 500.00
}
}`
	json2 := `{"kuo": "chaoYi"}`

	jb := new(database.BaseModel)
	jb.JsonStore.RawMessage = []byte(json2)

	db.AutoMigrate(&database.BaseModel{})
	db.Create(&database.BaseModel{
		ID:                  uuid.UUID{},
		CreatedAt:           time.Time{},
		UpdatedAt:           nil,
		BaseModelSoftDelete: database.BaseModelSoftDelete{},
		BaseModelJsonb:      database.BaseModelJsonb{
			JsonStore: postgres.Jsonb{
				RawMessage: []byte(json1),
			},
		},
	})

	db.Create(&jb)
	var result database.BaseModel
	db.First(&result)
	log.Printf("Println this a objcet: %s", &result)

	a := &result.JsonStore
	b, _ := json.Marshal(a)
	log.Printf("json.Marshal(): %s", a) // []byte
	os.Stdout.Write(b)                  // json

	/*	layout := "2006-01-02T15:04:05.000Z"
		str := &result.CreatedAt
		t, err := time.Parse(layout, str)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t)*/

	/*	if result.JsonStore.RawMessage == []byte(json2) {
			fmt.Println("SUCCESS: Selected JSON == inserted JSON")
		} else {
			fmt.Println("FAILED: Selected JSON != inserted JSON")
			fmt.Println("Inserted: " + STATE)
			fmt.Println("Selected: " + result.State)
		}*/
}
