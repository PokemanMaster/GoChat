package service

import (
	"IMProject/app/group/model"
	"IMProject/pkg/e"
	"IMProject/pkg/utils"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

type CreateGroupService struct {
	Name        string `json:"name" `       // 群组名称
	Description string `json:"description"` // 群组描述
	OwnerID     uint   `json:"owner_id"`    // 群主的用户ID
	AvatarURL   string `json:"avatar_url"`  // 群头像的URL
}

// Create 添加群聊
func (service *CreateGroupService) Create(c *gin.Context) *resp.Response {
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
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}

	// 返回数据
	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
