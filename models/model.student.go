package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelStudent struct {
	ID        string    `json:"id" gorm:"primary_key;"`
	Name      string    `json:"name,omitempty" gorm:"type:varchar(255);not null"`
	Npm       int       `json:"npm,omitempty" gorm:"type:bigint;unique;not null"`
	Fak       string    `json:"fak,omitempty" gorm:"type:varchar(255);not null"`
	Bid       string    `json:"bid,omitempty" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (model *ModelStudent) BeforeCreate(db *gorm.DB) error {
	model.ID = uuid.New().String()
	model.CreatedAt = time.Now().Local()
	return nil
}

func (model *ModelStudent) BeforeUpdate(db *gorm.DB) error {
	model.UpdatedAt = time.Now().Local()
	return nil
}
