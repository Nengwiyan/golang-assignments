package entity

import (
	"final-project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `gorm:"not null" json:"name" form:"name" valid:"required~Admin name is required"`
	Email     string    `gorm:"not null" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password  string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password must 6 characters or more"`
	Products  []Product `gorm:"contraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(a)

	if errCreate != nil {
		err = errCreate
		return
	}

	a.Password = helpers.HashPass(a.Password)

	err = nil
	return
}
