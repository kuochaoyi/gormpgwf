package main

import (
	"encoding/json"
	"fmt"
	"github.com/kuochaoyi/gormpgwf/utils"
	"log"
	"os"
	"time"

	"github.com/kuochaoyi/gormpgwf/database"
	_ "github.com/kuochaoyi/gormpgwf/database"

	// "github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kuochaoyi/chinese-calendar-golang/calendar"
)

type Demo struct {
	database.BaseModel
	database.BaseModelSoftDelete
	database.BaseModelJsonb
	database.BaseModelSerialID
}

func main() {
	// database.Open()
	// database.DBClient.Insert()

	db := database.DBClient.DB
	db.AutoMigrate(&Demo{})

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

	jb := new(Demo)
	jb.JsonStore.RawMessage = []byte(json2)

	serialNewly := new(utils.SerialPgHandler).SetSerial("demo")

	// db.AutoMigrate(&database.BaseModel{})
	db.Create(&Demo{
		BaseModel:           database.BaseModel{},
		BaseModelSoftDelete: database.BaseModelSoftDelete{},
		BaseModelJsonb: database.BaseModelJsonb{
			JsonStore: postgres.Jsonb{
				RawMessage: []byte(json1),
			},
		},
		BaseModelSerialID: database.BaseModelSerialID{
			SerialID: serialNewly,
		},
	})

	db.Create(&jb)
	var result Demo
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

	log.Printf("%s \n", new(utils.SerialPgHandler).SetSerial("base_models"))
	log.Printf("%s", new(utils.SerialPgHandler).Serial("base_models"))

	d := calendar.ByTimestamp(time.Now().Unix())

	// bytes, _ := c.ToJSON()
	bytes1, _ := d.ToJSON()

	// fmt.Println(string(bytes))
	fmt.Println(string(bytes1))

}
