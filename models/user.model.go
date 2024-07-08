package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/viniblima/atfilms/handlers"
	gorm "gorm.io/gorm"
)

/*
Modelo de usuário
*/
type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Name     string `json:"Name" validate:"required,min=3,max=32"`
	Email    string `gorm:"unique" json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=8"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.ID = uuid.NewV4().String()
	np, _ := handlers.HashPassword(user.Password)
	user.Password = np

	return
}
