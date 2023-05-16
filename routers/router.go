package routers

import (
	"github.com/EDDYCJY/go-gin-example/controller"
	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/gin-gonic/gin"

	_ "github.com/EDDYCJY/go-gin-example/docs"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	var userController = new(controller.UserController)
	noAuthAPI := r.Group("/api/v1")
	noAuthAPI.POST("/login", userController.Login)
	noAuthAPI.POST("/captcha", controller.Captcha)

	authApi := r.Group("/api/v1")
	// 使用jwt鉴权
	authApi.Use(jwt.JWT())
	{
		authApi.POST("/testAuth", controller.TestAuth)
		authApi.POST("/getUserList", userController.GetUserList)
		authApi.POST("/setUserEnable", userController.UpdateEnable)
		authApi.POST("/userUpdate", userController.Update)
		authApi.POST("/userCreate", userController.Create)
		authApi.POST("/userDelete", userController.Delete)
	}

	return r
}
