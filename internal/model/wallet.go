package model

import (
	"time"
)

type Wallet struct {
	MobileNumber int     `json:"mobile_number" gorm:"primarykey"`
	Balance      float64 `json:"balance"`
}

type Transaction struct {
	ID           int     `gorm:"primarykey"`
	MobileNumber int     `json:"mobile_number" gorm:"index"`
	Value        float64 `json:"value"`
	CreatedAt    time.Time
}
