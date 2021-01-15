package database

import (
	"encoding/json"
	"log"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// PostgreSQL - JSONB
type BaseModelJsonb struct {
	// State string `gorm:"type:JSONB NOT NULL DEFAULT '{}'"`
	JsonbStore postgres.Jsonb `gorm:"type:jsonb;default:'{}'"`
}

// Get a JSONB []byte. then: err = json.Unmarshal(v.([]byte), &s)
func (u *BaseModelJsonb) GetJsonbStore() (jsonBytes []byte) {
	v, err := u.JsonbStore.Value()
	if err != nil {
		log.Panicln(err)
	}

	if json.Valid(v.([]byte)) != false {
		log.Panicln(err)
	}

	/*	err = json.Unmarshal(v.([]byte), &pm)
		if err != nil {
			log.Panicln(err)
		}*/

	jb := v.([]byte)
	return jb
}

// Set a JSONB.
func (u *BaseModelJsonb) SetJsonbStore(JsonStruct interface{}) {
	b, err := json.Marshal(JsonStruct)
	if err != nil {
		log.Panicln(err)
	}
	u.JsonbStore = postgres.Jsonb{b}
}
