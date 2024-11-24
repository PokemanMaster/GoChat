package model

import (
	"github.com/PokemanMaster/GoChat/common/db"
	"gorm.io/gorm"
)

// Contact 人员关系
type Contact struct {
	gorm.Model
	OwnerID  uint //谁的关系信息
	TargetID uint //对应的谁 /群 ID
	Type     int  //对应的类型  1好友  2群
}

func (table *Contact) TableName() string {
	return "contact"
}

func SearchFriend(userId uint) []User {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	db.DB.Where("owner_id = ? and type=1", userId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetID))
	}
	users := make([]User, 0)
	db.DB.Where("id in ?", objIds).Find(&users)
	return users
}
