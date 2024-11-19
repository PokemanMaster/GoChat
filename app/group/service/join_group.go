package service

import (
	"IMProject/app/group/model"
	"IMProject/pkg/e"
	"IMProject/pkg/utils"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

type JoinGroupsService struct {
	UserId uint
	ComId  string
}

// Join  加入群聊
func (service *JoinGroupsService) Join(c *gin.Context) *resp.Response {
	userId := service.UserId
	comId := service.ComId
	data, msg := model.JoinGroup(userId, comId)
	if data == 0 {
		utils.RespOK(c.Writer, data, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
	code := e.SUCCESS
	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
