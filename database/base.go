package database

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

// BaseModel defines the common columns that all db structs should hold, usually
// db structs based on this have no soft delete
type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `gorm:"index"`
	BaseModelSoftDelete
	BaseModelJsonb
}

// BaseModelSoftDelete defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type BaseModelSoftDelete struct {
	// BaseModel
	DeletedAt *time.Time `sql:"index"`
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

// PostgreSQL - JSONB
type BaseModelJsonb struct {
	// State string `gorm:"type:JSONB NOT NULL DEFAULT '{}'"`
	JsonStore postgres.Jsonb `gorm:"type:jsonb"`
}
