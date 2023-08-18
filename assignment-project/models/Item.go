package models

import "time"

type Item struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	Description string     `gorm:"description"  json:"description"`
	Quantity    int        `gorm:"not null" json:"quantity"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	OrderID     uint
	Order       *Order
}
