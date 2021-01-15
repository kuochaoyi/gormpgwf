package main

import (
	"encoding/json"
	"log"

	"github.com/kuochaoyi/gormpgwf/database"
	_ "github.com/kuochaoyi/gormpgwf/database"
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
	db.AutoMigrate(&Demo{})

	db.Debug().AutoMigrate(&ClassRoom{})

	// JSON to insert
	STATE := `{"uses-kica": false, "hide-assessments-intro": true, "most-recent-grade-skew": 1.5}`

	classRoom := ClassRoom{State: STATE}
	db.Save(&classRoom)

	// Select the row
	var result ClassRoom
	db.First(&result)

	if result.State == STATE {
		log.Println("SUCCESS: Selected JSON == inserted JSON")
	} else {
		log.Println("FAILED: Selected JSON != inserted JSON")
		log.Println("Inserted: " + STATE)
		log.Println("Selected: " + result.State)
	}

	json3 := []byte(`{"age":  27, "name": "Yan"}`)

	a := &j{}
	err := json.Unmarshal(json3, &a)
	if err != nil {
		log.Fatalln("")
	}

	/*	json1 := []byte(`{
		"name": "Pasta",
			"ingredients": ["Flour", "Eggs", "Salt", "Water"],
				"organic": true,
				"dimensions": {
				"weight": 500.00
				}
			}`)

		json2 := []byte(`{"kuo": "chaoYi"}`)*/

	//jb := new(Demo)
	jb2 := &Demo{}
	jb2.SetJsonbStore(&a)

	// serialNewly := new(utils.SerialPgHandler).SetSerial("demo")

	// db.AutoMigrate(&database.BaseModel{})

	db.Create(&jb2)
	/*
		var result Demo
		db.First(&result)
		log.Printf("Println this a objcet: %s", &result)

		a := &result.
		b, _ := json.Marshal(a)
		log.Printf("json.Marshal(): %s", a) // []byte
		os.Stdout.Write(b)                  // json*/

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

	/*	log.Printf("%s \n", new(utils.SerialPgHandler).SetSerial("base_models"))
		log.Printf("%s", new(utils.SerialPgHandler).Serial("base_models"))

		d := calendar.ByTimestamp(time.Now().Unix())

		// bytes, _ := c.ToJSON()
		bytes1, _ := d.ToJSON()

		// fmt.Println(string(bytes))
		fmt.Println(string(bytes1))*/

}
