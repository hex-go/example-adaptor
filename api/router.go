package api

import (
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("/api/v1")
	InitUserRouter(ApiGroup) // 注册用户路由
	return Router
}
