package services

import (
	"errors"
	"strconv"
	"time"
	"xiaozhuquan.com/xiaozhuquan/app/common/request"
	"xiaozhuquan.com/xiaozhuquan/app/models"
	"xiaozhuquan.com/xiaozhuquan/global"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	/*var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error*/
	return
}

// Login 登录
func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	/*err = global.App.DB.Where("mobile = ?", params.Mobile).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}*/
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

func (userService *userService) GetSyncUserById(id string) (err error, syncUser models.SyncUser) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&syncUser, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

// 更新登录时间
func (userService *userService) UpdateLoginTime(userId string) error {
	intId, _ := strconv.Atoi(userId)
	return global.App.DB.Model(&models.SyncUser{}).Where("id = ?", intId).Update("login_time", time.Now()).Error
}

// 更新版本信息
func (userService *userService) UpdateSyncUserVersionById(userId string, version string) error {
	intId, _ := strconv.Atoi(userId)
	return global.App.DB.Model(&models.SyncUser{}).Where("id = ?", intId).Update("version", version).Error
}

// 通过ID获取User信息
func (userService *userService) GetUserById(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return err, user
}

// 通过unquie获取User信息
func (userService *userService) GetUserByUnquie(unique string) (err error, user models.User) {
	err = global.App.DB.First(&user, unique).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return err, user
}
