package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Variant struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UUID        string `json:"uuid"`
	VariantName string `gorm:"not null" json:"variant_name" form:"variant_name" valid:"required~Variant name is required, minstringlength(3)~Variant name must be 3 characters or more"`
	Quantity    int    `gorm:"not null" json:"quantity" form:"quantity" valid:"required~Quantity of product is required, numeric~Just number value allowed"`
	ProductID   uint   `form:"product_id"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(v)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
