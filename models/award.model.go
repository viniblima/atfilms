package models

import (
	uuid "github.com/satori/go.uuid"
	gorm "gorm.io/gorm"
)

type Award struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	Name       string `json:"Name" validate:"required,min=3,max=32"`
	AwardImage *Photo `gorm:"foreignKey:AwardImageID"`
	Jobs       []*Job `gorm:"many2many:job_awards;"`
}

func (m *Award) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}
