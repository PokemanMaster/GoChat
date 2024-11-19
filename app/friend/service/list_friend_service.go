package service

import (
	"IMProject/app/user/model"
	"IMProject/pkg/e"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

type FriendListsService struct {
	UserId uint
}

// List 获取好友列表
func (service *FriendListsService) List(c *gin.Context) *resp.Response {
	id := service.UserId
	users := model.SearchFriend(id)
	code := e.SUCCESS
	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   users,
	}
}
