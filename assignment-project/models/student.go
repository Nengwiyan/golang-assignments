package models

import "time"

type Student struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"not null" json:"studentName" form:"name"`
	Age       int        `gorm:"not null" json:"age" form:"age"`
	CreatedAt *time.Time `json:"createdAt" form:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" form:"updated_at"`
	Scores    []Score    `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;" json:"scores"`
}
