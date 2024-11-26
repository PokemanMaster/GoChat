package service

import (
	model2 "github.com/PokemanMaster/GoChat/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
	"go.uber.org/zap"
)

type FriendListsService struct {
}

// List 获取好友列表
func (service *FriendListsService) List(id string) *resp.Response {
	contacts := make([]model2.Contact, 0)
	err := db.DB.Model(&model2.Contact{}).Where("owner_id = ? and type=1", id).Find(&contacts).Error
	if err != nil {
		zap.L().Info("查询好友列表失败", zap.String("app.friend.service.list_friend", err.Error()))
		return &resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	list := make([]uint, 0)
	for _, contact := range contacts {
		list = append(list, contact.TargetID)
	}

	users := make([]model2.User, 0)
	err = db.DB.Model(model2.User{}).Where("id in ?", list).Find(&users).Error
	if err != nil {
		zap.L().Info("查询用户列表失败", zap.String("app.friend.service.list_friend", err.Error()))
		return &resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	return &resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
		Data:   users,
	}
}
