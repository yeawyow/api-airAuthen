package model

type User struct {
	Id int 
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
	Pname string `json:"Pname" binding:"required"`
	Fullname string `json:"Fullname" binding:"required"`
	DepartmentId int `json:"DepartmentId" binding:"required"`
	PositionId int `json:"PositionId" binding:"required"`
	StatusId string `json:"StatusId" binding:"required"`
	Avatar string `json:"Avatar" binding:"required"`
	Sex string `json:"Sex" bining:"required"`
}