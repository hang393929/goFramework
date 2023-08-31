package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
	"xiaozhuquan.com/xiaozhuquan/app/common/response"
	"xiaozhuquan.com/xiaozhuquan/app/models"
	"xiaozhuquan.com/xiaozhuquan/app/services"
	"xiaozhuquan.com/xiaozhuquan/global"
)

func CallBackAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		// 判断是否在checkAuthorityWhite中
		if checkAuthoritySpeciale("checkAuthorityWhite", path) {
			return
		}

		// 头信息是否完整
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.Fail(c, global.Errors.BadGatewayError.ErrorCode, "请登录")
			c.Abort()
			return
		}

		// redis中是否存在
		userId, err := global.App.Redis.Get(c.Request.Context(), tokenStr).Result()
		if err != nil || userId == "999999999999" {
			// 未查询到该用户信息
			response.Fail(c, global.Errors.BadGatewayError.ErrorCode, "未查询到该用户信息")
			c.Abort()
			return
		}

		// 验证用户是否属于平台
		err, syncUser := services.UserService.GetSyncUserById(userId)
		if err != nil {
			// 登录失效
			response.Fail(c, global.Errors.BadGatewayError.ErrorCode, "未查询到该用户信息")
			c.Abort()
			return
		}

		// 验证登录时间
		if !CheckLoginTime(syncUser) {
			// 登录失效
			response.Fail(c, global.Errors.BadGatewayError.ErrorCode, "未查询到该用户信息")
			c.Abort()
			return
		}

		// 更新登录时间并刷新token
		if !checkAuthoritySpeciale("checkAuthorityNotSetToken", path) && services.UserService.UpdateLoginTime(userId) != nil {
			global.App.Redis.Set(c, tokenStr, userId, time.Duration(60*60*8)*time.Second)
		}

		// 刷新版本
		if path == "checklogin" {
			version := c.GetHeader("use-version")
			if version == "" {
				version = "0"
			}

			if err := services.UserService.UpdateSyncUserVersionById(userId, version); err != nil {
				// 处理更新错误
				response.Fail(c, global.Errors.BadGatewayError.ErrorCode, "服务器异常")
				c.Abort()
				return
			}
		}
	}
}

func checkAuthoritySpeciale(key string, requestPath string) bool {
	// 加载配置文件
	viper.SetConfigFile("yaml/checkAuthorityWhite.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to read config file: " + err.Error())
	}

	checkAuthorityWhite := viper.GetStringSlice(key)

	// 匹配白名单
	for _, path := range checkAuthorityWhite {
		if path == requestPath {
			return true
		}
	}

	return false
}

func CheckLoginTime(syncUser models.SyncUser) bool {
	expire := time.Duration(8) * time.Hour
	loginTime := time.Unix(int64(syncUser.LoginTime), 0)
	currentTime := time.Now()

	if currentTime.Sub(loginTime) > expire {
		return false
	}

	return true
}
