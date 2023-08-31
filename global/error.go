package global

import "fmt"

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("%d: %s", c.ErrorCode, c.ErrorMsg)
}

type CustomErrors struct {
	BusinessError       CustomError
	ValidateError       CustomError
	TokenError          CustomError
	BadGatewayError     CustomError
	UserNotFoundError   CustomError
	DataExceptionError  CustomError
	WriteTaskFaildError CustomError
}

var Errors = CustomErrors{
	BusinessError:       CustomError{40000, "业务错误"},
	ValidateError:       CustomError{42200, "请求参数错误"},
	TokenError:          CustomError{40100, "登录授权失效"},
	BadGatewayError:     CustomError{502, "服务器异常"},
	UserNotFoundError:   CustomError{10001, "用户不存在"},
	DataExceptionError:  CustomError{10005, "数据异常"},
	WriteTaskFaildError: CustomError{11000, "写入任务失败"},
}
