package main

import (
	"github.com/cool-ops/gin-demo/controller"
	"github.com/cool-ops/gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) *gin.Engine{
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(),controller.Info)
	return r
}
