package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName      string    `gorm:"type:varchar(200);not null;unique;comment:'用户名'" json:"user_name"`
	Password      string    `gorm:"type:varchar(2000);not null;comment:'密码(AES加密)'" json:"password"`
	Telephone     string    `gorm:"type:char(11);comment:'手机号'" json:"telephone"`
	LevelID       uint      `gorm:"type:int unsigned;comment:'会员等级ID'" json:"level_id"`
	Avatar        string    `gorm:"type:varchar(200);comment:'用户头像'" json:"avatar"`
	Money         uint      `gorm:"type:int;comment:'用户金额';index:idx_money" json:"money"`
	Email         string    `gorm:"type:varchar(200);comment:'邮箱'" json:"email"`
	ClientIp      string    `gorm:"type:varchar(100);comment:'客户端IP'" json:"client_ip"`
	ClientPort    string    `gorm:"type:varchar(10);comment:'客户端端口'" json:"client_port"`
	Salt          string    `gorm:"type:varchar(200);comment:'加密盐'" json:"salt"`
	LoginTime     time.Time `gorm:"comment:'登录时间';default:null" json:"login_time"`
	HeartbeatTime time.Time `gorm:"comment:'心跳时间';default:null" json:"heartbeat_time"`
	LoginOutTime  time.Time `gorm:"comment:'登出时间';column:login_out_time;default:null" json:"login_out_time"`
	IsLogout      bool      `gorm:"comment:'是否登出'" json:"is_logout"`
	DeviceInfo    string    `gorm:"type:varchar(500);comment:'设备信息'" json:"device_info"`
}
