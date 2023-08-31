package models

import "strconv"

// 这里有的用了指针，是为了区别是真的为空还是为0，指针类型的零值是nil，而不是0或空字符串，
// 对于数值类型（如int、float等）和字符串类型，它们的零值是0和空字符串""，这和它们的实际值可能是相同的
type User struct {
	ID
	Account          string  `json:"account"`
	Platform         int     `json:"platform"`
	UserName         string  `json:"user_name"`
	Project          uint8   `json:"project"`
	Classify         uint8   `json:"classify"`
	Channel          int     `json:"channel"`
	UserID           *int    `json:"user_id"`
	Unique           string  `json:"unique"`
	SyncID           *int    `json:"sync_id"`
	Field            string  `json:"field"`
	Original         uint8   `json:"original"`
	Level            *int    `json:"level"`
	Status           int     `json:"status"`
	Quality          string  `json:"quality"`
	Credit           string  `json:"credit"`
	CreditStr        string  `json:"credit_str"`
	Fens             int     `json:"fens"`
	MCNID            *int    `json:"mcn_id"`
	Video            *int    `json:"video"`
	Message          string  `json:"message"`
	LevelName        string  `json:"level_name"`
	GCardCount       string  `json:"gcard_count"`
	TotalPlay        *int64  `json:"total_play"`
	UpdateTime       *int    `json:"update_time"`
	BiliScore        *int    `json:"bili_score"`
	Cookie           string  `json:"cookie"`
	CookieStatus     int     `json:"cookie_status"`
	KBStatus         int     `json:"kb_status"`
	CreateTime       *int    `json:"create_time"`
	UpdateCookieTime *int    `json:"update_cookie_time"`
	Page             string  `json:"page"`
	Like             *int    `json:"like"`
	Heads            string  `json:"heads"`
	Head             string  `json:"head"`
	CookieText       string  `json:"cookie_text"`
	Follow           *int    `json:"follow"`
	FinderUsername   string  `json:"finder_username"`
	Unique2          *string `json:"unique_2"`
	NickName         string  `json:"nick_name"`
	Unique3          string  `json:"unique_3"`
	Phone            string  `json:"phone"`
	Link             string  `json:"link"`
	ChannelChild     uint8   `json:"channel_child"`
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
