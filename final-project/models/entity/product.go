package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UUID      string `json:"uuid"`
	Name      string `gorm:"not null" json:"name" form:"name" valid:"required~Product name is required, minstringlength(3)~Product name must be 3 characters or more"`
	ImageURL  string `json:"image_url"`
	Variants  []Variant
	AdminID   uint
	Admin     *Admin
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
