package controller

import (
	"fmt"
	"github.com/cool-ops/gin-demo/common"
	"github.com/cool-ops/gin-demo/dto"
	"github.com/cool-ops/gin-demo/model"
	"github.com/cool-ops/gin-demo/response"
	"github.com/cool-ops/gin-demo/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": dto.ToUserDTO(user.(model.User)),
		},
	})
}

// 用户注册
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
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号不能为空或者必须是11位")
		return
	}

	if len(password) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能为空")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	if len(name) == 0 {
		name = utils.RandomString(10)
		fmt.Println(name)
	}
	// 数据库中查找手机号是否存在，如果存在，则返回已注册
	if isTelephoneExist(db, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号码已被注册")
		return
	}

	// 密码加密
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "密码加密失败")
		return
	}
	// 开始注册
	newUser := model.User{
		UserName:  name,
		PassWord:  string(hasePassword),
		Telephone: telephone,
	}
	db.Create(&newUser)

	// 生成token
	token, err := common.GenerateToken(newUser)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code": 500,
			"msg":  "系统错误",
		})
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统错误")
		log.Println("generate token failed. err : " + err.Error())
		return
	}
	// 注册成功
	response.Success(ctx, gin.H{"token": token,}, "注册成功")
}

// 用户登录
func Login(ctx *gin.Context) {
	// 获取登录所需的信息，用户名（可选），手机号，密码
	username := ctx.PostForm("username")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 判断输入的合法性
	if len(username) == 0 && len(telephone) == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "请输入正确的用户名")
		return
	}
	if len(telephone) < 11 {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "手机号输入有误，必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码不符合规范，必须多于6位")
		return
	}
	// 到数据库中进行校验
	var db = common.GetDB()
	var user model.User
	var dbName = viper.GetString("db.dbName")
	if err := db.Model(dbName).Where("telephone = ?", telephone).First(&user).Error; err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "手机号有误，请检查")
		log.Println(err.Error())
		return
	}
	if user.ID == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "用户名或密码错误，请检查")
		return
	}
	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password));err != nil{
		response.Response(ctx, http.StatusBadRequest, 400, nil, "用户名或密码错误，请检查")
		return
	}
	// 校验成功返回token
	token, err := common.GenerateToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常，请稍后再试")
		return
	}
	response.Success(ctx,gin.H{"token":token,},"登录成功")
}

// 判断手机号是否已经注册
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
