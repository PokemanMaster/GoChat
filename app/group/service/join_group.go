package service

import (
	"IMProject/app/user/model"
	"IMProject/common/db"
	"IMProject/pkg/e"
	"IMProject/resp"
)

type JoinGroupService struct {
	UserID  uint
	GroupID uint
}

// Join  加入群聊
func (service *JoinGroupService) Join() *resp.Response {
	// 获取数据
	UserID := service.UserID
	GroupID := service.GroupID

	// 存储数据
	var contact model.Contact
	contact.OwnerID = UserID
	contact.TargetID = GroupID
	contact.Type = 2

	// 逻辑处理
	err := db.DB.Model(&contact).Create(&contact).Error
	if err != nil {
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 返回数据
	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
