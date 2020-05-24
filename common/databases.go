package common

import (
	"fmt"
	"github.com/cool-ops/gin-demo/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB {
	driverName := viper.GetString("db.dbDriver")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	dbName := viper.GetString("db.dbName")
	charSet := viper.GetString("db.charSet")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host,port,dbName, charSet,
	)
	db, err := gorm.Open(driverName, args)
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
