package models

import (
	uuid "github.com/satori/go.uuid"
	gorm "gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	Name       string `json:"Name" validate:"required,min=3,max=32"`
	Logo       *Photo `json:"Logo" gorm:"foreignKey:LogoID"`
	ShowInHome *bool  `json:"ShowInHome" validate:"required"`
	Position   int    `json:"Position" validate:"required"`
}

func (m *Customer) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}
