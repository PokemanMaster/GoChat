package service

import (
	"github.com/PokemanMaster/GoChat/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
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
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	// 返回数据
	return &resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
	}
}
