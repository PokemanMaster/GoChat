package service

import (
	"IMProject/app/user/model"
	"IMProject/common/db"
	"IMProject/pkg/e"
	"IMProject/resp"
)

type GroupFriendListsService struct {
	TargetId uint
}

// List 获取群聊里的好友列表
func (service *GroupFriendListsService) List() resp.Response {
	targetId := service.TargetId

	var contact []model.Contact
	db.DB.Where("target_id=? and type=2", targetId).Find(&contact)

	code := e.SUCCESS
	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   contact,
	}
}
