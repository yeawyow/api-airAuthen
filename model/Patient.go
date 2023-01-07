package model

import (
	"time"
)

type Patient struct {
	HosGuid  string
	Hn       string `gorm:"primaryKey"`
	Pname    string
	Fname    string `json:"fname" binding:"required"`
	Lname    string
	Birthday time.Time
	Cid string `json:"cid"`
	Hometel string `gorm:"home_tel"`


}