package common

import (
	"fmt"
	"github.com/cool-ops/gin-demo/model"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	username := "root"
	password := "coolops@123456"
	dbName := "user"
	charSet := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host,port,dbName, charSet,
	)
	db, err := gorm.Open("mysql", args)
	if err != nil {
		log.Println("connect to MySQL failed. err " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB()*gorm.DB{
	return DB
}
