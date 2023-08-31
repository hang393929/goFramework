package services

import (
	"xiaozhuquan.com/xiaozhuquan/app/models"
	"xiaozhuquan.com/xiaozhuquan/global"
)

type videoService struct {
}

var VideoService = new(videoService)

// 新增 videoData 记录
func (videoService *videoService) AddVideoData(createTime int, date string, userID int, content string) error {
	videoData := models.VideoData{
		CreateTime: createTime,
		Date:       date,
		UserID:     userID,
		Content:    content,
	}

	if err := global.App.DB.Create(&videoData).Error; err != nil {
		return global.CustomError{ErrorCode: 11000, ErrorMsg: "写入任务失败"}
	}

	return nil
}
