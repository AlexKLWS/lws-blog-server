package models

import "time"

// Based off gorm.Model
type Model struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func (m Model) GetID() uint {
	return m.ID
}

func (m Model) GetCreatedAt() time.Time {
	return m.CreatedAt
}
