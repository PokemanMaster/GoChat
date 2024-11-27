package serializer

import (
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
