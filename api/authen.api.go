package api

import (
	"main/db"
	"main/interceptor"
	"main/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(passwordjwt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordjwt), 14)
	return string(bytes), err
}
func setupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api")
	{
      
		authenAPI.POST("/login", login)
		//authenAPI.GET("/login", test)
		//authenAPI.POST("/register", register)
	}
}

type LoginBody struct{
	Username string `json:"Username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func login(c *gin.Context) {
	var json LoginBody
	if c.ShouldBindJSON(&json) == nil {
		var queryUser model.User
		if err := db.GetDB().First(&queryUser, "username=?", json.Username).Error; err != nil {
			c.JSON(200, gin.H{"result": "nok", "error": err})
		} else if CheckPasswordHash(json.Password, queryUser.Password) == false {
			c.JSON(200, gin.H{"result": "nok"})
		} else {

			token := interceptor.JwtSign(queryUser)

			c.JSON(200, gin.H{"result": "ok", "accessToken": token,"username":queryUser.Username,"roles":"ROLE_MODERATOR","uid":queryUser.Id})
		}

	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}

}

/*func register(c *gin.Context) {
	var PsMember model.PsMember

	if c.ShouldBind(&PsMember) == nil {
		PsMember.Passwordjwt, _ = hashPassword(PsMember.Passwordjwt)
		if err := db.GetDB().Create(&PsMember).Error; err != nil {
			c.JSON(200, gin.H{"status": "nok", "error": err})
		} else {
			c.JSON(200, gin.H{"result": "ok", "data": PsMember})
		}

	} else {
		c.JSON(200, gin.H{"data": "unable to bind data"})
	}

}*/
