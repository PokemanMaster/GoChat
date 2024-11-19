package service

import (
	"IMProject/app/group/model"
	"IMProject/pkg/e"
	"IMProject/pkg/utils"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

type CreateGroupService struct {
	OwnerId uint
	Name    string
	Icon    string
	Desc    string
}

// Create 添加群聊
func (service *CreateGroupService) Create(c *gin.Context) *resp.Response {
	ownerId := service.OwnerId
	name := service.Name
	icon := service.Icon
	desc := service.Desc

	community := model.GroupBasic{}
	community.OwnerId = ownerId
	community.Name = name
	community.Img = icon
	community.Desc = desc

	code, msg := model.CreateCommunity(community)
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}

	cod := e.SUCCESS
	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(cod),
	}
}
