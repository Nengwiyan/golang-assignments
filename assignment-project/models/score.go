package models

import "time"

type Score struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	AssignmentTitle string     `gorm:"not null" json:"assignment_title"`
	Description     string     `gorm:"description"  json:"description"`
	Scores          int        `gorm:"not null" json:"scores"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
	StudentID       uint
	Student         *Student
}
