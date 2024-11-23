package model

import (
	"IMProject/common/db"
	utils2 "IMProject/pkg/mid"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

// GetUserList 获取用户列表
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	db.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	db.DB.Where("name = ? and pass_word=?", name, password).First(&user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils2.MD5Encode(str)
	db.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	db.DB.Where("name = ?", name).First(&user)
	return user
}
func CreateUser(user UserBasic) *gorm.DB {
	return db.DB.Create(&user)
}
func DeleteUser(user UserBasic) *gorm.DB {
	return db.DB.Delete(&user)
}

// FindByID 查找某个用户
func FindByID(id uint) UserBasic {
	user := UserBasic{}
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
