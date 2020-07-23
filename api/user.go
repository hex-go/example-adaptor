package api

import (
	"example/adaptor"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/users")
	//UserRouter.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.POST("", CreateUser)       // 创建用户
		UserRouter.GET("/policies", Policies) // 设置用户租户
	}
}

type Name struct {
	Name string `json:"name"`
}

// @Tags KeyUser
// @Summary 创建用户
// @Produce  application/json
// @Param data body model.SysUser true "用户创建接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var req Name
	_ = c.ShouldBindJSON(&req)

	for _, plugin := range adaptor.FactoryByName {
		status, pErr := plugin().CreateUser(req.Name)
		if status != true {
			c.JSON(500, gin.H{
				"message": pErr,
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Success!",
		})
		return
	}
}

// @Tags KeyUser
// @Summary 规则查询
// @Produce  application/json
// @Param data body model.SysUser true "规则查询接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Router /users/policies [get]
func Policies(c *gin.Context) {
	for _, plugin := range adaptor.FactoryByName {
		status, pErr := plugin().Policies()
		if status != true {
			c.JSON(500, gin.H{
				"message": pErr,
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Success!",
		})
		return
	}
}
