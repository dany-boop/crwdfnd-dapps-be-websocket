package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID    uint
	Content   string
	CreatedAt time.Time
}
