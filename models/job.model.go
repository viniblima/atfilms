package models

import (
	uuid "github.com/satori/go.uuid"
	gorm "gorm.io/gorm"
)

type Job struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	Name       string `json:"Name" validate:"required,min=3,max=32"`
	Slug       string `json:"Slug" validate:"required,min=3,max=32"`
	CustomerID string `json:"CustomerID" validate:"required"`
	// JobID       string          `json:"JobID" validate:"required"`
	Customer    *Customer       `gorm:"foreignKey:CustomerID"`
	ShowInHome  bool            `json:"ShowInHome" validate:"required"`
	Components  *[]JobComponent `gorm:"foreignKey:JobID"`
	Placeholder Photo           `gorm:"foreignKey:PlaceholderID"`
	MainVideo   Video           `gorm:"foreignKey:MainVideoID"`
	Tags        []*Tag          `gorm:"many2many:job_tags;"`
}

func (m *Job) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}
