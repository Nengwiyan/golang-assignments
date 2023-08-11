package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Variants struct {
	ID          uint   `gorm:"PrimaryKey"`
	VariantName string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	ProductID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (v *Variants) BeforeCreate(tx *gorm.DB) (err error) {

	if len(v.VariantName) < 3 {
		err = errors.New("Variant name is too short. Minimal variant name is 3")
	}
	return
}
