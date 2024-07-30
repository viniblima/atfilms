package models

import (
	uuid "github.com/satori/go.uuid"
	gorm "gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"unique" json:"Name" validate:"required,min=3,max=32"`
	Jobs []*Job `gorm:"many2many:job_tags;"`
}

func (m *Tag) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}
