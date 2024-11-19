package service

import (
	"IMProject/app/group/model"
	"IMProject/pkg/e"
	"IMProject/pkg/utils"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

type JoinGroupService struct {
	UserID  uint
	GroupID uint
}

// Join  加入群聊
func (service *JoinGroupService) Join(c *gin.Context) *resp.Response {
	// 获取数据
	UserID := service.UserID
	GroupID := service.GroupID

	// 逻辑处理
	data, msg := model.JoinGroup(UserID, GroupID)
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
