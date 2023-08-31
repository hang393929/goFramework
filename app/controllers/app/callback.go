package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xiaozhuquan.com/xiaozhuquan/app/common/response"
	"xiaozhuquan.com/xiaozhuquan/app/services"
	"xiaozhuquan.com/xiaozhuquan/global"
)

type CallBackController struct{}

/**
*回调用户
 */
func (call *CallBackController) User(c *gin.Context) {
	userIdStr := c.DefaultPostForm("user_id", "0")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		response.FailByError(c, global.Errors.UserNotFoundError)
		return
	}

	dataStr := c.PostForm("data")
	if dataStr == "" {
		response.FailByError(c, global.Errors.DataExceptionError)
		return
	}

	err = services.CallBackService.HaddleCallBackUser(dataStr, userId)
	if err != nil {
		info := global.CustomError{ErrorCode: 11000, ErrorMsg: "写入任务失败"}
		response.FailByError(c, info)
		return
	}

	response.Success(c, []interface{}{})
}

/**
*回调视频
 */
func (call *CallBackController) Video(c *gin.Context) {
	dataStr := c.PostForm("data")
	if dataStr == "" {
		response.FailByError(c, global.Errors.DataExceptionError)
		return
	}

	err := services.CallBackService.HaddleCallBackVideo(dataStr)
	if err.ErrorCode != 0 {
		info := global.CustomError{ErrorCode: 11000, ErrorMsg: "写入任务失败"}
		response.FailByError(c, info)
		return
	}

	response.Success(c, []interface{}{})
}

/**
* 回调粉丝
 */
func (call *CallBackController) Fens(c *gin.Context) {
	dataStr := c.PostForm("data")
	if dataStr == "" {
		response.FailByError(c, global.Errors.DataExceptionError)
		return
	}

	err := services.CallBackService.HaddleCallBackFens(dataStr)
	if err != nil {
		info := global.CustomError{ErrorCode: 11000, ErrorMsg: "写入任务失败"}
		response.FailByError(c, info)
		return
	}

	response.Success(c, []interface{}{})
}

/**
* 粉丝（旧）
 */
func (call *CallBackController) Income(c *gin.Context) {
	dataStr := c.PostForm("data")
	if dataStr == "" {
		response.FailByError(c, global.Errors.DataExceptionError)
		return
	}

	err := services.CallBackService.HaddleCallBackIncome(dataStr)
	if err != nil {
		info := global.CustomError{ErrorCode: 11000, ErrorMsg: "写入任务失败"}
		response.FailByError(c, info)
		return
	}

	response.Success(c, []interface{}{})
}

/**
* 回调data
 */
func (call *CallBackController) Data(c *gin.Context) {
	dataStr := c.PostForm("data")
	if dataStr == "" {
		response.FailByError(c, global.Errors.DataExceptionError)
		return
	}

	err := services.CallBackService.HaddleCallBackData(dataStr)
	if err != nil {
		info := global.CustomError{ErrorCode: 11000, ErrorMsg: "写入任务失败"}
		response.FailByError(c, info)
		return
	}

	response.Success(c, []interface{}{})
}
