package models

import (
	uuid "github.com/satori/go.uuid"
	gorm "gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	Name string `json:"Name" validate:"required,min=3,max=32"`
}

func (m *Client) BeforeCreate(db *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}
