package models

import (
	"time"

	"gorm.io/gorm"
)

type Donation struct {
	gorm.Model
	UserID    uint
	Amount    float64 `gorm:"not null"`
	Message   string
	ChartLink string
	VideoLink string
	CreatedAt time.Time
}
