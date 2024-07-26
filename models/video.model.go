package models

import (
	uuid "github.com/satori/go.uuid"
	gorm "gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ID                   string `gorm:"primaryKey"`
	MainVideoID          *string
	JobComponentVideosID *string
	FileName             string `json:"FileName" validate:"required"`
}

func (m *Video) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}
