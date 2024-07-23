package models

import (
	uuid "github.com/satori/go.uuid"
	gorm "gorm.io/gorm"
)

type JobComponent struct {
	gorm.Model
	ID                  string           `gorm:"primaryKey"`
	JobID               string           `json:"JobID" validate:"required"`
	Type                jobComponentType `json:"Type" validate:"required"`
	Title               string
	Text                string
	Slider              []Photo `gorm:"foreignKey:ReferenceID"`
	Videos              []Video `gorm:"foreignKey:ReferenceID"`
	FillPhotoHorizontal Photo   `gorm:"foreignKey:ReferenceID"`
	Position            int     `json:"Position" validate:"required"`
}

func (m *JobComponent) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}

type CreateJobComponentStruct struct {
	Type     jobComponentType `json:"Type" validate:"required"`
	Position int              `json:"Position" validate:"required"`
}
