package model

import (
	"github.com/PokemanMaster/GoChat/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/common/db"
	"gorm.io/gorm"
)

// GroupBasic 群模块
type GroupBasic struct {
	gorm.Model
	Name           string `json:"name" gorm:"Index:idx_name"`    // 群组名称
	Description    string `json:"description"`                   // 群组描述
	OwnerID        uint   `json:"owner_id"`                      // 群主的用户ID
	AvatarURL      string `json:"avatar_url"`                    // 群头像的URL
	MaxMembers     uint   `json:"max_members" gorm:"default:60"` // 群组成员上限
	CurrentMembers uint   `json:"current_members"`               // 当前群成员数
	IsPublic       bool   `json:"is_public" gorm:"default:true"` // 是否公开群组
	IsActive       bool   `json:"is_active" gorm:"default:true"` // 群组是否可用
}

// CreateGroup 创建群
func CreateGroup(group GroupBasic) (status int, message string) {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if len(group.Name) == 0 {
		return 400, "群名称不能为空"
	}
	if group.OwnerID <= 0 {
		return 400, "请先登录"
	}
	if err := db.DB.Create(&group).Error; err != nil {
		tx.Rollback()
		return 500, "创建群失败"
	}

	contact := model.Contact{}
	contact.OwnerID = group.OwnerID
	contact.TargetID = group.ID
	contact.Type = 2
	if err := db.DB.Create(&contact).Error; err != nil {
		tx.Rollback()
		return 500, "添加群关系失败"
	}

	tx.Commit()
	return 200, "创建成功"
}
