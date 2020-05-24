package middleware

import (
	"fmt"
	"github.com/cool-ops/gin-demo/common"
	"github.com/cool-ops/gin-demo/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// 用户认证的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求中获取Authorization
		tokenString := ctx.GetHeader("Authorization")
		fmt.Println(tokenString)
		// 判断Authorization是否合法
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer"){
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg": "权限不足",
			})
			ctx.Abort()
			return
		}
		log.Println("Authorization 合法")
		// 解析Authorization
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg": "权限不足",
			})
			ctx.Abort()
			return
		}
		log.Println("token 解析成功")
		// 获取user信息
		userId := claims.UserID
		db := common.GetDB()
		var user model.User
		db.First(&user,userId)

		// 用户不存在，返回401
		if user.ID == 0{
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg": "权限不足",
			})
			ctx.Abort()
			return
		}
		// 用户存在，写入上下文
		ctx.Set("user",user)
		ctx.Next()
	}
}
