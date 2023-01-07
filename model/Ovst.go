package model

type Ovst struct {
	HosGuid string `json:"hos_guid" binding:"required"`
	Vn      string `json:"vn" binding:"required"`
	Hn      string `json:"hn" binding:"required"`
	Vstdate string `json:"vstdate" binding:"required"`
	Vsttime string `json:"vsttime" binding:"required"`

}