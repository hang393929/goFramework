package common

import (
	"github.com/gin-gonic/gin"
	"xiaozhuquan.com/xiaozhuquan/app/common/request"
	"xiaozhuquan.com/xiaozhuquan/app/common/response"
	"xiaozhuquan.com/xiaozhuquan/app/services"
)

func ImageUpload(c *gin.Context) {
	var form request.ImageUpload
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	outPut, err := services.MediaService.SaveImage(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, outPut)
}
