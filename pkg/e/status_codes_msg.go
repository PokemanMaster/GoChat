package e

// MsgFlags 状态码map
var MsgFlags = map[int]string{
	SUCCESS:      "成功",
	ERROR_PARAMS: "参数无效",
	ERROR:        "标准性错误",

	ERROR_USERNAME_EXIST:      "用户名已存在",
	ERROR_ACCOUNT_EXIST:       "用户已存在",
	ERROR_ACCOUNT_NOT_EXIST:   "用户不存在",
	ERROR_MATCHED_USERNAME:    "用户名必须为 6-12 位英文和数字",
	ERROR_ACCOUNT_NOT_COMPARE: "账号只允许数字或字母",
	ERROR_PASSWORD:            "密码错误",
	ERROR_UPDATE_PASSWORD:     "密码更新错误",
	ERROR_ENCRYPTION_FAIL:     "密码加密失败",
	ERROR_ACCOUNT_NOT_EMPTY:   "账号不能为空",
	ERROR_PASSWORD_NOT_EMPTY:  "密码不能为空",
	ERROR_MATCHED_PASSWORD:    "密码长度6-12位且包含字母或数字",
	ERROR_CAPTCHA:             "验证码错误",
	ERROR_PASSWORD_CONFIRM:    "前后密码不一致",

	ERROR_NOT_EXIST_PRODUCT: "产品不存在",
	ERROR_NOT_EXIST_ADDRESS: "地址不存在",
	ERROR_EXIST_FAVORITE:    "收藏已存在",
	ERROR_NOT_EXIST_FRIEND:  "好友不存在",
	ERROR_CODE:              "验证码错误",

	ERROR_AUTH_TOKEN_FAIL:    "用户认证失败",
	ERROR_AUTH_TOKEN_TIMEOUT: "用户认证超时",
	ERROR_CREATED_ACCOUNT:    "创建用户失败",
	ERROR_UNMARSHAL_JSON:     "反序列化JSON失败",
	ERROR_DATABASE:           "数据库操作错误",
	ERROR_OSS:                "对象存储服务错误",
	ERROR_WEBSOCKET_CREATE:   "websocket 创建错误",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
