package controller

import (
	"fmt"
	"github.com/cool-ops/gin-demo/common"
	"github.com/cool-ops/gin-demo/model"
	"github.com/cool-ops/gin-demo/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	// 获取注册数据
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	telephone := ctx.PostForm("telephone")
	// 校验用户名、密码、手机号码
	// 手机号码必须是11位，如果手机号存在则返回已注册
	// 密码不能为空
	// 用户名如果为空，则生成十位随机字符串作为用户名

	if len(telephone) != 11 || len(telephone) == 0 {
		ctx.JSON(422, gin.H{
			"code": 422,
			"msg":  "手机号不能为空或者必须是11位",
		})
		return
	}

	if len(password) == 0 {
		ctx.JSON(422, gin.H{
			"code": 422,
			"msg":  "密码不能为空",
		})
		return
	}

	if len(name) == 0 {
		name = utils.RandomString(10)
		fmt.Println(name)
	}
	// 数据库中查找手机号是否存在，如果存在，则返回已注册
	if isTelephoneExist(db,telephone){
		ctx.JSON(422,gin.H{
			"code":422,
			"msg":"手机号码已被注册.",
		})
		return
	}

	// 开始注册
	newUser := model.User{
		UserName:  name,
		PassWord:  password,
		Telephone: telephone,
	}
	db.Create(&newUser)
	// 注册成功
	ctx.JSON(200, gin.H{
		"code":      200,
		"name":      name,
		"password":  password,
		"telephone": telephone,
		"msg":       "注册成功",
	})
}


// 判断手机号是否已经注册
func isTelephoneExist(db *gorm.DB,telephone string)bool{
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return  true
	}
	return false
}