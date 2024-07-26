package models

import (
	uuid "github.com/satori/go.uuid"
	gorm "gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID                    string  `gorm:"primaryKey"`
	LogoID                *string `json:"LogoID" `
	PlaceholderID         *string
	SliderID              *string
	FillPhotoHorizontalID *string
	FileName              string `json:"FileName" validate:"required"`
}

func (m *Photo) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}
