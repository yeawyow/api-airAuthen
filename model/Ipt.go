package model

import "time"

type Ipt struct {
	An string `json:"an" binding:"required"`
	Hn string
	//Dchdate time.Time
	//Dchtype string
	Pttype   string
	Ward     string
	Nameward string
	Ptname   string
	//Regdate string
	Fullname string
	Regtime  time.Time
}