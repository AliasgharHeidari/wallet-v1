package model

import (
	"time"
)

type Wallet struct {
	MobileNumber int `json:"mobileNumber" gorm:"primarykey"`
	Balance      int `json:"balace"`
	UpdatedAt    time.Time
}
