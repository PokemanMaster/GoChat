package service

import (
	"IMProject/app/group/model"
	"IMProject/pkg/e"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

type GroupListsService struct {
	OwnerId uint
}

// List 加载群列表
func (service *GroupListsService) List(c *gin.Context) *resp.Response {
	ownerId := service.OwnerId
	data, msg := model.LoadCommunity(ownerId)
	code := e.SUCCESS

	if len(data) != 0 {
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   data,
		}
	} else {
		return &resp.Response{
			Status: code,
			Msg:    msg,
		}
	}

}
