package models

import (
	uuid "github.com/satori/go.uuid"
	gorm "gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	Name       string `json:"Name" validate:"required,min=3,max=32"`
	Logo       Photo  `gorm:"foreignKey:ReferenceID"`
	ShowInHome *bool  `json:"ShowInHome"`
	Index      uint   `json:"Index" validate:"required" gorm:"unique;autoIncrement"`
}

func (m *Customer) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}
