package database

import (
	"testing"
	"time"

	"gorm.io/gorm"
)



import (
	"testing"
	"time"


	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestBaseModel_BeforeCreate(t *testing.T) {
	type fields struct {
		ID                  uuid.UUID
		CreatedAt           time.Time
		UpdatedAt           *time.Time
		BaseModelSoftDelete BaseModelSoftDelete
	}
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			fields:  fields{
				ID:                  uuid.UUID{},
				CreatedAt:           time.Time{},
				UpdatedAt:           nil,
				BaseModelSoftDelete: BaseModelSoftDelete{
					DeletedAt: gorm.DeletedAt{},},
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := &BaseModel{
				ID:                  tt.fields.ID,
				CreatedAt:           tt.fields.CreatedAt,
				UpdatedAt:           tt.fields.UpdatedAt,
				BaseModelSoftDelete: tt.fields.BaseModelSoftDelete,
			}
			if err := base.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
