package serializer

import (
	"IMProject/app/user/model"
	"time"
)

// UserSerialization 用户序列化器
type UserSerialization struct {
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Avatar        string
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLogout      bool
	DeviceInfo    string
}

// BuildUser 序列化用户
func BuildUser(user model.UserBasic) UserSerialization {
	return UserSerialization{

		Avatar: user.Avatar,
	}
}
