package db

import (
	"context"
	"fmt"
	"main/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"

	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB{
	return db
}


type SqlLogger struct {
  logger.Interface	
}
func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error){
	sql , _ :=fc() 
	fmt.Printf("%v\n==================\n",sql)
}



func SetupDB(){
//var Ipt []model.Ipt
//var Patient []model.Patient
var IptNhsoImage []model.IptNhsoImage
	
	dsnhos := os.Getenv("MYSQL_DNSH")
	database,err := gorm.Open(mysql.Open(dsnhos), &gorm.Config{
		Logger: &SqlLogger{},
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //nochang table name
		},
	})
	database.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsnhos)},
	//	Policy: dbresolver.RandomPolicy{},
		
	},&IptNhsoImage))
	
	

	if err != nil{
          panic("fail to connet databese")
	}
  
	database.AutoMigrate()
	
	db=database

}
