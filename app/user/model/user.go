package model

import (
	"fmt"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/mid"
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
	LoginTime     time.Time `gorm:"comment:'登录时间'" json:"login_time"`
	HeartbeatTime time.Time `gorm:"comment:'心跳时间'" json:"heartbeat_time"`
	LoginOutTime  time.Time `gorm:"comment:'登出时间';column:login_out_time" json:"login_out_time"`
	IsLogout      bool      `gorm:"comment:'是否登出'" json:"is_logout"`
	DeviceInfo    string    `gorm:"type:varchar(500);comment:'设备信息'" json:"device_info"`
}

// GetUserList 获取用户列表
func GetUserList() []*User {
	data := make([]*User, 10)
	db.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByNameAndPwd(name string, password string) User {
	user := User{}
	db.DB.Where("name = ? and pass_word=?", name, password).First(&user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := mid.MD5Encode(str)
	db.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func FindUserByName(name string) User {
	user := User{}
	db.DB.Where("name = ?", name).First(&user)
	return user
}
func CreateUser(user User) *gorm.DB {
	return db.DB.Create(&user)
}
func DeleteUser(user User) *gorm.DB {
	return db.DB.Delete(&user)
}

// FindByID 查找某个用户
func FindByID(id uint) User {
	user := User{}
	db.DB.Where("id = ?", id).First(&user)
	return user
}

// AddFriend 添加好友   自己的ID  ， 好友的ID
func AddFriend(userId uint, targetName string) (int, string) {
	if targetName != "" {
		targetUser := FindUserByName(targetName)
		//fmt.Println(targetUser, " userId        ", )
		if targetUser.Salt != "" {
			if targetUser.ID == userId {
				return -1, "不能加自己"
			}
			contact0 := Contact{}
			db.DB.Where("owner_id =?  and target_id =? and type=1", userId, targetUser.ID).Find(&contact0)
			if contact0.ID != 0 {
				return -1, "不能重复添加"
			}
			tx := db.DB.Begin()
			//事务一旦开始，不论什么异常最终都会 Rollback
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()
			contact := Contact{}
			contact.OwnerID = userId
			contact.TargetID = targetUser.ID
			contact.Type = 1
			if err := db.DB.Create(&contact).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}
			contact1 := Contact{}
			contact1.OwnerID = targetUser.ID
			contact1.TargetID = userId
			contact1.Type = 1
			if err := db.DB.Create(&contact1).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}
			tx.Commit()
			return 0, "添加好友成功"
		}
		return -1, "没有找到此用户"
	}
	return -1, "好友ID不能为空"
}
