package database

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel defines the common columns that all db structs should hold, usually
// db structs based on this have no soft delete
type BaseModel struct {
	// `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time  `gorm:"column:created_at;default:null" json:"created_at"` // (My|Postgre)SQL
	UpdatedAt time.Time `gorm:"column:updated_at;default:null" json:"updated_at"`
	BaseModelSoftDelete
}

type BaseModel1 struct {
	gorm.Model
	// `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
}

// BaseModelSoftDelete defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type BaseModelSoftDelete struct {
	// BaseModel
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;default:null" json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
// https://gorm.io/docs/create.html#Create-Hooks
func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	// NewRandom returns a Random (Version 4) UUID.
	id, err := uuid.NewRandom()
	log.Printf("UUID v4 genter to: %s", id)
	if err != nil {
		return err
	}

	tx.Statement.SetColumn("ID", id.String())
	return err
}

// Serial id = date + 6 digital ex. 20200123000001
type BaseModelSerialID struct {
	SerialID string `gorm:"type:varchar(14)"`
}
