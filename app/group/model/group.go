package model

import (
	"IMProject/app/user/model"
	"IMProject/common/db"
	"fmt"

	"gorm.io/gorm"
)

// GroupBasic 群模块
type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Img     string
	Desc    string
}

// CreateCommunity 创建群
func CreateCommunity(community GroupBasic) (int, string) {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if len(community.Name) == 0 {
		return -1, "群名称不能为空"
	}
	if community.OwnerId == 0 {
		return -1, "请先登录"
	}
	if err := db.DB.Create(&community).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return -1, "建群失败"
	}
	contact := model.Contact{}
	contact.OwnerId = community.OwnerId
	contact.TargetId = community.ID
	contact.Type = 2 //群关系
	if err := db.DB.Create(&contact).Error; err != nil {
		tx.Rollback()
		return -1, "添加群关系失败"
	}
	tx.Commit()
	return 0, "建群成功"
}

// LoadCommunity 查找群
func LoadCommunity(ownerId uint) ([]*GroupBasic, string) {
	contacts := make([]model.Contact, 0)
	objIds := make([]uint64, 0)
	db.DB.Where("owner_id = ? and type = 2", ownerId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}

	data := make([]*GroupBasic, 10)
	db.DB.Where("id in ?", objIds).Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}

	return data, "查询成功"
}

// JoinGroup 加入群
func JoinGroup(userId uint, comId string) (int, string) {
	contact := model.Contact{}
	contact.OwnerId = userId
	contact.Type = 2
	group := GroupBasic{}

	db.DB.Where("id=? or name=?", comId, comId).Find(&group)
	if group.Name == "" {
		return -1, "没有找到群"
	}
	db.DB.Where("owner_id=? and target_id=? and type = 2 ", userId, comId).Find(&contact)
	if !contact.CreatedAt.IsZero() {
		return -1, "已加过此群"
	} else {
		contact.TargetId = group.ID
		db.DB.Create(&contact)
		return 0, "加群成功"
	}
}
