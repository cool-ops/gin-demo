package main

import (
	"github.com/cool-ops/gin-demo/common"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = LoadRouter(r)

	panic(r.Run())
}




