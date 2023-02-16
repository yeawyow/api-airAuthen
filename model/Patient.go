package model

import (
	"time"
)

type Patient struct {
	HosGuid  string
	Hn       string `json:"hn" gorm:"primaryKey"`
	Pname    string
	Fname    string `json:"fname" binding:"required"`
	Lname    string
	Birthday time.Time
	Cid      string `json:"cid"`
	Hometel  string `gorm:"home_tel"`
	Hcode    string `gorm:"hcode"`
}
