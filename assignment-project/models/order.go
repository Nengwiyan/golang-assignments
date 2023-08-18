package models

import "time"

type Order struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	CustomerName string     `gorm:"not null" json:"customerName" form:"customerName"`
	OrderedAt    *time.Time `json:"orderedAt" form:"ordered_at"`
	CreatedAt    *time.Time `json:"createdAt" form:"created_at"`
	UpdatedAt    *time.Time `json:"updatedAt" form:"updated_at"`
	Items        []Item     `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;" json:"items"`
}
