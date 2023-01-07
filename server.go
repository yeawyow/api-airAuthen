package main

import (
	"fmt"
	"main/api"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
)

func main(){
 err := godotenv.Load(".env")
 if err !=nil {
	fmt.Println("Error loading .env file")
 }
	router := gin.Default()
	//setup CORS midleware Option
	config := cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: false,
		ValidateHeaders: false,
	}
	router.Use(cors.Middleware(config))
	api.Setup(router)
	router.Run(":8081")
}