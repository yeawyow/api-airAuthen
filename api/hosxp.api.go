package api

import (
	"fmt"
	"main/db"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func setupHosxpAPI(router *gin.Engine) {
	
	authenAPI := router.Group("/api")
	{

		
		//authenAPI.GET("/getmembers", getmembers)
		authenAPI.POST("/register", register)
		authenAPI.POST("/registvisit", regvisit)
		authenAPI.GET("/getpatient/:id", getpatient)
		authenAPI.GET("/gettest", gettest)
		authenAPI.GET("/getovst/:hn", getovst)
	
	}
}


func regvisit(c *gin.Context){
	//var ovst model.Ovst
	//db.GetDB()
	c.JSON(200,gin.H{"ok":"ok"})
}

func register(c *gin.Context){

	var json model.User
	if err := c.ShouldBindJSON(&json); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Chec user exists
	var userExist model.User
	db.GetDB().Where("username =?", json.Username).First(&userExist)
	if userExist.Id > 0{
		c.JSON(http.StatusOK, gin.H{"status": "error","message":"User create failed"})
	}else{
		encryptedPassword , _ := bcrypt.GenerateFromPassword([]byte(json.Password),14)
	user := model.User{Username: json.Username,Password: string(encryptedPassword),
	Pname: json.Pname,Fullname: json.Fullname,DepartmentId: json.DepartmentId,
	PositionId: json.PositionId,StatusId: json.StatusId,Avatar: json.Avatar,}
if result := db.GetDB().Create(&user); result.Error != nil{
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": result.Error.Error(),
	})
	return
}
	
c.JSON(http.StatusCreated, &user)
	}
	
	
	
}
func getpatient(c *gin.Context){
	var Patient model.Patient
	id := c.Param("id")
	
	
if 	tx:=db.GetDB().Where("cid =?",id).First(&Patient).Error; tx != nil{
   c.JSON(200,gin.H{"result":"","error":tx})
 // c.AbortWithStatus(http.StatusNotFound)
} else{

	c.JSON(http.StatusOK,Patient)
	//c.JSON(http.StatusOK,gin.H{"result":Patient})
	/*c.JSON(200, gin.H{
		"patient":Patient,
	})*/
}

	
	
}
func getovst(c *gin.Context){
	var ovst model.Ovst
	id := c.Param("hn")
	//vstdate := time.Now()
	
if 	tx:=db.GetDB().Where("hn =?",id).Find(&ovst).Error; tx != nil{
   // c.JSON(200,gin.H{"result":"","error":tx})
   c.AbortWithStatus(http.StatusNotFound)
} else{
	c.JSON(http.StatusOK,ovst)
	/*c.JSON(200, gin.H{
		"patient":Patient,
	})*/
}

	
	
}


func gettest(c *gin.Context){
	var IptNhsoImage []model.IptNhsoImage
	//var Ipt []model.Ipt
	tx:=db.GetDB().Last(&IptNhsoImage)
	if tx.Error !=nil{
		fmt.Println(tx.Error)
		return
	}
	c.JSON(200,IptNhsoImage)
}

