package models

import "time"

type Product struct {
	ID        uint   `gorm:"PrimaryKey"`
	Name      string `gorm:"not null; type varchar(115)"`
	Variants  []Variants
	CreatedAt time.Time
	UpdatedAt time.Time
}
