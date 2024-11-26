package service

import (
	"github.com/PokemanMaster/GoChat/server/app/group/model"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
)

type CreateGroupService struct {
	Name        string `json:"name" `       // 群组名称
	Description string `json:"description"` // 群组描述
	OwnerID     uint   `json:"owner_id"`    // 群主的用户ID
	AvatarURL   string `json:"avatar_url"`  // 群头像的URL
}

// Create 添加群聊
func (service *CreateGroupService) Create() *resp.Response {
	// 获取数据
	name := service.Name
	ownerId := service.OwnerID
	avatar := service.AvatarURL
	desc := service.Description

	// 存储数据
	group := model.GroupBasic{}
	group.Name = name
	group.OwnerID = ownerId
	group.AvatarURL = avatar
	group.Description = desc

	// 逻辑处理
	code, msg := model.CreateGroup(group)
	if code != 200 {
		return &resp.Response{
			Status: code,
			Msg:    msg,
		}
	}

	// 返回数据
	return &resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
	}
}
