package service

import (
	"github.com/PokemanMaster/GoChat/app/user/model"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
)

type DeleteFriendService struct {
	UserID   uint
	FriendID uint
}

// Delete 删除好友
func (service *DeleteFriendService) Delete() *resp.Response {
	// 检查好友关系是否存在
	var contact model.Contact
	err := db.DB.Model(&model.Contact{}).
		Where("owner_id = ? AND target_id = ? AND type = ?", service.UserID, service.FriendID, 1).
		First(&contact).Error
	if err != nil {
		return &resp.Response{
			Status: e.ERROR_NOT_EXIST_FRIEND,
			Msg:    e.GetMsg(e.ERROR_NOT_EXIST_FRIEND),
		}
	}

	// 删除好友关系
	tx := db.DB.Begin()

	err = tx.Delete(&contact).Error
	if err != nil {
		tx.Rollback()
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 双向删除好友
	err = tx.Where("owner_id = ? AND target_id = ? AND type = ?", service.FriendID, service.UserID, 1).
		Delete(&model.Contact{}).Error
	if err != nil {
		tx.Rollback()
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	tx.Commit()

	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
