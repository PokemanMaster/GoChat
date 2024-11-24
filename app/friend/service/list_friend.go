package service

import (
	"github.com/PokemanMaster/GoChat/app/user/model"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
)

type FriendListsService struct {
}

// List 获取好友列表
func (service *FriendListsService) List(id string) *resp.Response {
	contacts := make([]model.Contact, 0)

	err := db.DB.Model(&model.Contact{}).Where("owner_id = ? and type=1", id).Find(&contacts).Error
	if err != nil {
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	list := make([]uint, 0)
	for _, contact := range contacts {
		list = append(list, contact.TargetID)
	}

	users := make([]model.UserBasic, 0)
	err = db.DB.Model(model.UserBasic{}).Where("id in ?", list).Find(&users).Error
	if err != nil {
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data:   users,
	}
}
