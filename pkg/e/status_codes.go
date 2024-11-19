package e

const (
	SUCCESS      = 200 // 标准性成功
	ERROR_PARAMS = 400 // 参数无效
	ERROR        = 500 // 标准性错误

	ERROR_USERNAME_EXIST      = 10001 // 用户名已存在
	ERROR_ACCOUNT_EXIST       = 10002 // 用户已存在
	ERROR_ACCOUNT_NOT_EXIST   = 10003 // 用户不存在
	ERROR_MATCHED_USERNAME    = 10004 // 用户名必须为 6-12 位英文和数字
	ERROR_ACCOUNT_NOT_COMPARE = 10005 // 账号只允许数字或字母
	ERROR_PASSWORD            = 10006 // 密码错误
	ERROR_UPDATE_PASSWORD     = 10007 // 密码更新错误
	ERROR_ENCRYPTION_FAIL     = 10008 // 密码加密失败
	ERROR_ACCOUNT_NOT_EMPTY   = 10009 // 账号不能为空
	ERROR_PASSWORD_NOT_EMPTY  = 10010 // 密码不能为空
	ERROR_MATCHED_PASSWORD    = 10011 // 密码长度6-12位且包含字母或数字
	ERROR_CAPTCHA             = 10012 // 验证码错误
	ERROR_PASSWORD_CONFIRM    = 10013 // 前后密码不一致

	ERROR_NOT_EXIST_PRODUCT = 20001 // 产品不存在
	ERROR_NOT_EXIST_ADDRESS = 20002 // 地址不存在
	ERROR_EXIST_FAVORITE    = 20003 // 收藏已存在
	ERROR_NOT_EXIST_FRIEND  = 20004 // 好友不存在

	ERROR_AUTH_TOKEN_FAIL    = 50001 // 用户认证失败
	ERROR_AUTH_TOKEN_TIMEOUT = 50002 // 用户认证超时
	ERROR_CREATED_ACCOUNT    = 50003 // 创建用户失败
	ERROR_UNMARSHAL_JSON     = 50004 // 反序列化JSON失败
	ERROR_DATABASE           = 50005 // 数据库操作错误
	ERROR_OSS                = 50006 // 对象存储服务错误
	ERROR_WEBSOCKET_CREATE   = 50007 // websocket 创建错误
)
