package models

import (
	uuid "github.com/satori/go.uuid"
	gorm "gorm.io/gorm"
)

type Job struct {
	gorm.Model
	ID          string         `gorm:"primaryKey"`
	Name        string         `json:"Name" validate:"required,min=3,max=32"`
	Slug        string         `json:"Slug" validate:"required,min=3,max=32"`
	CustomerID  string         `json:"CustomerID" validate:"required"`
	Customer    Customer       `gorm:"foreignKey:CustomerID" json:"Customer"`
	ShowInHome  bool           `json:"ShowInHome" validate:"required"`
	Components  []JobComponent `gorm:"foreignKey:JobID" `
	Placeholder Photo          `gorm:"foreignKey:PlaceholderID"`
	Video       Video          `gorm:"foreignKey:ReferenceID"`
}

func (m *Job) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}
