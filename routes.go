package main

import (
	"github.com/gin-gonic/gin"
	"junsonjack.cn/go_vue/controller"
	"junsonjack.cn/go_vue/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine{
	r.POST("/api/auth/register",controller.Register)
	r.POST("/api/auth/login",controller.Login)
	//使用中间件保护info接口
	r.GET("/api/auth/info",middleware.AuthMiddleware(),controller.Info)

	return r
}

