package models

type VideoData struct {
	ID         int    `json:"id" gorm:"column:id"`
	CreateTime int    `json:"create_time" gorm:"column:create_time"`
	Date       string `json:"date" gorm:"column:date"`
	Content    string `json:"content" gorm:"column:content"`
	UserID     int    `json:"user_id" gorm:"column:user_id"`
}
