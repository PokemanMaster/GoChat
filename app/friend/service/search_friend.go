package service

import (
	"IMProject/app/user/model"
	"IMProject/common/db"
	"IMProject/pkg/e"
	"IMProject/resp"
)

type SearchFriendService struct {
	UserID uint // 用户id
	Name   uint // 好友名字
}

// Search 查找好友
func (service *SearchFriendService) Search() *resp.Response {
	UserID := service.UserID

	contacts := make([]model.Contact, 0)
	list := make([]uint, 0)

	err := db.DB.Model(&model.Contact{}).Where("owner_id = ? and type=1", UserID).Find(&contacts).Error
	if err != nil {
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

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
	}
}