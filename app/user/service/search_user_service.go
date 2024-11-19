package service

import (
	"IMProject/app/user/model"
	"IMProject/pkg/e"
	"IMProject/pkg/utils"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

type SearchFriendService struct {
	UserId uint
}

// Search 查找某个用户
func (service *SearchFriendService) Search(c *gin.Context) *resp.Response {
	userId := service.UserId
	data := model.FindByID(userId)
	utils.RespOK(c.Writer, data, "ok")
	code := e.SUCCESS
	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
