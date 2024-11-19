package service

import (
	"IMProject/app/user/model"
	"IMProject/pkg/e"
	"IMProject/pkg/utils"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

type CreateFriendService struct {
	UserId     uint   // 自己的id
	TargetName string // 好友的id
}

// Create 添加用户
func (service *CreateFriendService) Create(c *gin.Context) *resp.Response {
	userId := service.UserId
	targetName := service.TargetName
	code, msg := model.AddFriend(userId, targetName)
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}

	cod := e.SUCCESS
	return &resp.Response{
		Status: cod,
		Msg:    e.GetMsg(cod),
	}
}
