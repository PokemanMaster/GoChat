package service

import (
	"errors"
	"github.com/PokemanMaster/GoChat/v1/server/app/user/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CreateFriendService struct {
	UserID   uint
	TargetID uint
}

func (service *CreateFriendService) Create() *resp.Response {
	UserID := service.UserID
	TargetID := service.TargetID

	// 检查是否已经是好友
	var exist model.Contact
	err := db.DB.Model(&model.Contact{}).
		Where("owner_id = ? AND target_id = ? AND type = ?", UserID, TargetID, 1).
		First(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Info("可以添加好友", zap.String("app.friend.service.create_friend", ""))
		} else {
			zap.L().Error("查询用户失败", zap.String("app.friend.service.create_friend", err.Error()))
			return &resp.Response{
				Status: e.ERROR_DATABASE,
				Msg:    e.GetMsg(e.ERROR_DATABASE),
				Error:  err.Error(),
			}
		}
	} else {
		zap.L().Info("好友已存在", zap.String("app.user.service.user_register", ""))
		return &resp.Response{
			Status: e.ERROR_ALREADY_FRIENDS,
			Msg:    e.GetMsg(e.ERROR_ALREADY_FRIENDS),
		}
	}

	tx := db.DB.Begin()

	// 创建好友关系
	var contact model.Contact
	contact.OwnerID = UserID
	contact.TargetID = TargetID
	contact.Type = 1

	err = tx.Create(&contact).Error
	if err != nil {
		tx.Rollback()
		zap.L().Error("创建用户失败", zap.String("app.friend.service.create_friend", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	contact.ID = 0
	contact.OwnerID = TargetID
	contact.TargetID = UserID
	contact.Type = 1

	err = tx.Create(&contact).Error
	if err != nil {
		tx.Rollback()
		zap.L().Error("创建双向好友失败", zap.String("app.friend.service.create_friend", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	tx.Commit()

	// 添加成功
	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
