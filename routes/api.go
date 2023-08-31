package routes

import (
	"github.com/gin-gonic/gin"
	"xiaozhuquan.com/xiaozhuquan/app/controllers/app"
	"xiaozhuquan.com/xiaozhuquan/app/controllers/common"
	"xiaozhuquan.com/xiaozhuquan/app/middleware"
	"xiaozhuquan.com/xiaozhuquan/app/services"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}

	callbackauthRouter := router.Group("").Use(middleware.CallBackAuth())
	{
		callbackauthRouter.POST("/callbackUser", (&app.CallBackController{}).User)     // 回调解析用户
		callbackauthRouter.POST("/callbackVideo", (&app.CallBackController{}).Video)   // 回调解析视频
		callbackauthRouter.POST("/callbackFens", (&app.CallBackController{}).Fens)     // 回调解析粉丝
		callbackauthRouter.POST("/callbackIncome", (&app.CallBackController{}).Income) // 回调解析粉丝
		callbackauthRouter.POST("/callbackData", (&app.CallBackController{}).Data)     // 回调解析粉丝
	}
}
