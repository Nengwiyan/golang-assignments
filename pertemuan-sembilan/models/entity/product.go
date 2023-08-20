package entity

import "time"

type Product struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	Description string     `gorm:"not null" json:"description"`
	Price       int64      `gorm:"not null" json:"price"`
	Stock       int        `gorm:"not null" json:"stock"`
	ImageURL    string     `json:"image_url"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
