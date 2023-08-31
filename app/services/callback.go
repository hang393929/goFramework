package services

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"xiaozhuquan.com/xiaozhuquan/app/models"
	"xiaozhuquan.com/xiaozhuquan/global"
)

type callBackService struct {
}

var CallBackService = new(callBackService)

/**
*用户入队
 */
func (call *callBackService) HaddleCallBackUser(dataStr string, userId int) (err global.CustomError) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
		return global.Errors.DataExceptionError
	}

	if userId == 0 {
		userId = int(data["user_id"].(float64))
	} else {
		data["user_id"] = userId
	}

	// 过滤重复数据
	key := fmt.Sprintf("CallbackUser_%d", userId) // 根据用户ID生成对应的缓存键名
	selectDB := 3                                 // 切换 Redis 数据库的选择
	if !call.repeatFilter(data, key, selectDB) {
		return
	}

	// 用户是否存在于平台
	_, exist := call.checkUserIsExistById(userId, false)
	if exist == false {
		return global.Errors.UserNotFoundError
	}

	// 写入到Redis的第2个数据库
	if err := call.writeToRedis(data, "callBackUserForGo", 3); err != nil {
		// 处理返回的错误
		return global.Errors.WriteTaskFaildError
	}

	return
}

/**
*视频入队
 */
func (call *callBackService) HaddleCallBackVideo(dataStr string) (err global.CustomError) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
		return global.Errors.DataExceptionError
	}

	if data["user_id"] == "" {
		return global.Errors.UserNotFoundError
	}

	var userID int
	if userIDValue, ok := data["user_id"].(int); ok {
		// 如果 data["user_id"] 是 int 类型，直接使用
		userID = userIDValue
	} else if userIDStr, ok := data["user_id"].(string); ok {
		// 尝试将 data["user_id"] 转换为字符串
		parsedUserID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return global.Errors.DataExceptionError
		}
		// 转换成功，将整数值赋给 userID
		userID = parsedUserID
	} else {
		// 处理 data["user_id"] 不是 int 或 string 类型的情况
		return global.Errors.DataExceptionError
	}

	_, exist := call.checkUserIsExistById(userID, false)
	if exist == false {
		return global.Errors.UserNotFoundError
	}

	if _, res := call.lockForCallBackVideo(userID, true); res == true {
		return
	}

	// 写入redis第2个数据库
	if err := call.writeToRedis(data, "callBackVideoForGo", 2); err != nil {
		// 处理返回的错误
		return global.Errors.WriteTaskFaildError
	}

	// 写入video_data
	videoService := &videoService{}
	createTime := int(time.Now().Unix())
	date := time.Now().Format("2006-01-02")
	content, _ := json.Marshal(data)
	dataJsonStr := string(content)
	if err := videoService.AddVideoData(createTime, date, userID, dataJsonStr); err != nil {
		// 处理错误
		return global.CustomError{ErrorCode: 11000, ErrorMsg: "写入任务失败"}
	}

	// 刷新锁
	_, _ = call.lockForCallBackVideo(userID, false)

	return
}

/**
* 粉丝入队
 */
func (call *callBackService) HaddleCallBackFens(dataStr string) error {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
		return global.Errors.DataExceptionError
	}

	if data["unique"] == "" {
		return global.Errors.UserNotFoundError
	}

	var unquie string
	if unquieValue, ok := data["unique"].(string); ok {
		unquie = unquieValue
	} else {
		return global.Errors.DataExceptionError
	}

	// 用户是否存在
	err, user := UserService.GetUserByUnquie(unquie)
	if err != nil {
		return global.Errors.UserNotFoundError
	}

	// 过滤重复数据
	key := fmt.Sprintf("FensRepeatFilter_%d", user.ID) // 根据用户ID生成对应的缓存键名
	selectDB := 4                                      // 切换 Redis 数据库的选择
	if !call.repeatFilter(data, key, selectDB) {
		return nil
	}

	// 写入到Redis的第4个数据库
	if err := call.writeToRedis(data, "callBackFensForGo", 4); err != nil {
		// 处理返回的错误
		return global.Errors.WriteTaskFaildError
	}

	return nil
}

/**
* income入队
 */
func (call *callBackService) HaddleCallBackIncome(dataStr string) error {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
		return global.Errors.DataExceptionError
	}

	if data["user_id"] == "" {
		return global.Errors.UserNotFoundError
	}

	var userID int
	if userIDValue, ok := data["user_id"].(int); ok {
		// 如果 data["user_id"] 是 int 类型，直接使用
		userID = userIDValue
	} else if userIDStr, ok := data["user_id"].(string); ok {
		// 尝试将 data["user_id"] 转换为字符串
		parsedUserID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return global.Errors.DataExceptionError
		}
		// 转换成功，将整数值赋给 userID
		userID = parsedUserID
	} else {
		// 处理 data["user_id"] 不是 int 或 string 类型的情况
		return global.Errors.DataExceptionError
	}

	user, exist := call.checkUserIsExistById(userID, true)
	if exist == false {
		return global.Errors.UserNotFoundError
	}

	// 过滤重复数据
	key := fmt.Sprintf("FensRepeatFilter_%d", user.ID) // 根据用户ID生成对应的缓存键名
	selectDB := 1                                      // 切换 Redis 数据库的选择
	if !call.repeatFilter(data, key, selectDB) {
		return nil
	}

	// 写入到Redis的第4个数据库
	if err := call.writeToRedis(data, "callBackFensForGo", 1); err != nil {
		// 处理返回的错误
		return global.Errors.WriteTaskFaildError
	}

	return nil
}

/**
* data入队
 */
func (call *callBackService) HaddleCallBackData(dataStr string) error {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
		return global.Errors.DataExceptionError
	}

	if data["user_id"] == "" {
		return global.Errors.UserNotFoundError
	}

	var userID int
	if userIDValue, ok := data["user_id"].(int); ok {
		// 如果 data["user_id"] 是 int 类型，直接使用
		userID = userIDValue
	} else if userIDStr, ok := data["user_id"].(string); ok {
		// 尝试将 data["user_id"] 转换为字符串
		parsedUserID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return global.Errors.DataExceptionError
		}
		// 转换成功，将整数值赋给 userID
		userID = parsedUserID
	} else {
		// 处理 data["user_id"] 不是 int 或 string 类型的情况
		return global.Errors.DataExceptionError
	}

	user, exist := call.checkUserIsExistById(userID, true)
	if exist == false {
		return global.Errors.UserNotFoundError
	}

	// 过滤重复数据
	key := fmt.Sprintf("CallbackData_%d", user.ID) // 根据用户ID生成对应的缓存键名
	selectDB := 1                                  // 切换 Redis 数据库的选择
	if !call.repeatFilter(data, key, selectDB) {
		return nil
	}

	// 写入到Redis的第4个数据库
	if err := call.writeToRedis(data, "callBackFensForGo", 1); err != nil {
		// 处理返回的错误
		return global.Errors.WriteTaskFaildError
	}

	return nil
}

// repeatFilter 实现幂等性校验，将 $data 存入缓存并设置过期时间为 1 小时
func (call *callBackService) repeatFilter(data map[string]interface{}, key string, selectDB int) bool {
	ctx := context.Background()

	// 切换 Redis 数据库
	global.App.Redis.Do(ctx, "SELECT", selectDB)

	cachedData, err := global.App.Redis.Get(ctx, key).Result()
	if err == nil {
		// 缓存存在，判断是否与 $data 相同
		if cachedData == call.getMD5Hash(data) {
			return false
		}
	}

	// 缓存不存在，将 $data 存入缓存
	global.App.Redis.Set(ctx, key, call.getMD5Hash(data), time.Hour)

	return true
}

// 计算 MD5 哈希值
func (call *callBackService) getMD5Hash(data map[string]interface{}) string {
	dataStr, _ := json.Marshal(data)

	// 创建 MD5 哈希算法的实例
	hasher := md5.New()

	// 将数据写入哈希实例
	hasher.Write(dataStr)

	// 计算哈希值并返回为字节数组
	hashBytes := hasher.Sum(nil)

	// 将字节数组转换为十六进制字符串
	hashStr := fmt.Sprintf("%x", hashBytes)

	return hashStr
}

// 检测用户是否存在,并根据情况是否返回user对象(这里返回指针，可以在后续直接修改属性的值，返回对象只能使用数据，不能修改)
func (call *callBackService) checkUserIsExistById(userId int, returnUser bool) (*models.User, bool) {
	// 这里是根据用户ID查询用户信息的代码逻辑，您需要根据您的数据库操作来实现此部分
	err, user := UserService.GetUserById(strconv.Itoa(userId))
	if err != nil {
		return nil, false
	} else {
		// 判断是否查询到了有效的用户数据
		if user.GetUid() == "0" {
			return nil, false
		}
	}

	// 进一步检查用户的同步信息，您可以根据您的业务逻辑来实现此部分
	err, syncUser := UserService.GetSyncUserById(strconv.Itoa(*user.UserID))
	if err != nil || syncUser.ID == 0 {
		return nil, false
	}

	// 如果 returnUser 为 true，则返回用户对象
	if returnUser {
		return &user, true
	}

	// 否则，只返回 true 表示用户存在
	return nil, true
}

// 写入redis
func (call *callBackService) writeToRedis(data map[string]interface{}, key string, selectDb int) error {
	// 创建一个上下文
	ctx := context.Background()

	// 切换到Redis的第三个数据库
	global.App.Redis.Do(ctx, "SELECT", selectDb)

	// 将数据转换为JSON字符串
	dataStr := call.toJsonString(data)

	// 将数据写入Redis
	_, err := global.App.Redis.Set(ctx, key, dataStr, 0).Result()
	if err != nil {
		// 处理写入失败的情况
		return err
	}

	return nil
}

func (call *callBackService) toJsonString(data map[string]interface{}) string {
	dataStr, _ := json.Marshal(data)
	return string(dataStr)
}

// lockForCallBackVideo 回调视频数据锁
func (call *callBackService) lockForCallBackVideo(userId int, get bool) (error, bool) {
	fiveMinuteKey := fmt.Sprintf("callVideo_fiveMinute_%d", userId)

	// 创建一个上下文
	ctx := context.Background()

	if get == true {
		callVideoKey := fmt.Sprintf("callVideo_%d", userId)
		callVideo, err := global.App.Redis.Get(ctx, callVideoKey).Result()
		if err != nil || callVideo == "" {
			return nil, false
		}

		fiveMinute, err := global.App.Redis.Get(ctx, callVideoKey).Result()
		if err != nil || fiveMinute != "" {
			return nil, false
		}

		return nil, true
	}

	// 设置锁
	callVideoKey := fmt.Sprintf("callVideo_%d", userId)
	if _, err := global.App.Redis.Get(ctx, callVideoKey).Result(); err != nil {
		userIdStr := strconv.Itoa(userId)
		// 如果不存在，则设置锁并设置过期时间
		global.App.Redis.Set(ctx, callVideoKey, userIdStr, time.Hour)
		global.App.Redis.Set(ctx, fiveMinuteKey, userIdStr, 5*time.Minute)
		return nil, true
	}

	return nil, false
}
