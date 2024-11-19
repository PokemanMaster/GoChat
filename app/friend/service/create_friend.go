package service

import (
	"IMProject/app/user/model"
	"IMProject/common/db"
	"IMProject/pkg/e"
	"IMProject/resp"
)

type CreateFriendService struct {
	UserID   uint
	TargetID uint
}

// Create 添加用户
func (service *CreateFriendService) Create() *resp.Response {
	UserID := service.UserID
	TargetID := service.TargetID

	var contact model.Contact
	contact.OwnerID = UserID
	contact.TargetID = TargetID
	contact.Type = 1

	err := db.DB.Model(&contact).Create(&contact).Error
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
