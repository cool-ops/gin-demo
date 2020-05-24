package main

import (
	"github.com/cool-ops/gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) *gin.Engine{
	r.POST("/api/auth/register", controller.Register)
	return r
}
