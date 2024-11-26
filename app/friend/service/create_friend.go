package service

import (
	"errors"
	"github.com/PokemanMaster/GoChat/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
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
				Status: e2.ERROR_DATABASE,
				Msg:    e2.GetMsg(e2.ERROR_DATABASE),
				Error:  err.Error(),
			}
		}
	} else {
		zap.L().Info("好友已存在", zap.String("app.user.service.user_register", ""))
		return &resp.Response{
			Status: e2.ERROR_ALREADY_FRIENDS,
			Msg:    e2.GetMsg(e2.ERROR_ALREADY_FRIENDS),
		}
	}

	// 创建好友关系
	var contact model.Contact
	contact.OwnerID = UserID
	contact.TargetID = TargetID
	contact.Type = 1

	err = db.DB.Create(&contact).Error
	if err != nil {
		zap.L().Error("创建用户失败", zap.String("app.friend.service.create_friend", err.Error()))
		return &resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	// 添加成功
	return &resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
	}
}
