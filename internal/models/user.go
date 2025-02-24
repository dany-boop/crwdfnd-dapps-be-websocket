package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	Admin  Role = "admin"
	Member Role = "user"
)

type User struct {
	gorm.Model
	Username   string `gorm:"unique;not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	WalletAddr string `gorm:"unique;not null"`
	Role       Role   `gorm:"type:user_role;default:'user'"` // ENUM Role
}
