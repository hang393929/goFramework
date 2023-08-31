package models

type SyncUser struct {
	ID           int
	Account      string `json:"account" gorm:"size:64;not null;comment:账号"`
	Channel      int    `json:"channel" gorm:"type:tinyint(2);not null;default:1;comment:渠道 1 影视 2 体育 3 数科"`
	SyncID       int    `json:"sync_id" gorm:"not null;comment:原平台id"`
	Project      int    `json:"project" gorm:"type:tinyint(3) unsigned;not null;default:1;comment:1小猪优版 2小猪视觉"`
	Classify     int    `json:"classify" gorm:"type:tinyint(3) unsigned;not null;default:1;comment:1泛娱乐 2泛生活"`
	ChannelChild int    `json:"channel_child" gorm:"type:int(10) unsigned;not null;default:0;comment:小猪优版帐号子分类 泛娱乐下面有：1影视2体育3游戏 ， 泛生活下面有 1：3c数码 2美妆 3母婴 4百货"`
	Mcn          int    `json:"mcn" gorm:"type:tinyint(2);not null;default:2;comment:1 是 2 不是"`
	McnID        int    `json:"mcn_id" gorm:"not null;default:0"`
	CreateTime   int    `json:"create_time" gorm:"not null;comment:创建时间"`
	Status       int    `json:"status" gorm:"type:tinyint(2);not null;default:1;comment:1 正常 2 锁定"`
	UserName     string `json:"user_name" gorm:"size:255"`
	Heads        string `json:"heads" gorm:"size:255;default:''"`
	Head         string `json:"head" gorm:"size:255"`
	UpdateTime   int    `json:"update_time"`
	LoginTime    int    `json:"login_time" gorm:"not null;default:0;comment:登录时间"`
	Token        string `json:"token" gorm:"size:255"`
	Version      string `json:"version" gorm:"size:256;not null;default:'';comment:版本号"`
	AreaCode     int64  `json:"area_code" gorm:"type:bigint(20) unsigned;not null;default:0;comment:运营地区码"`
}
