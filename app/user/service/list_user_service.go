package service

import (
	"IMProject/app/user/model"
	"IMProject/pkg/e"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

type UserListsService struct {
}

// List
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func (service *UserListsService) List(c *gin.Context) *resp.Response {
	data := make([]*model.UserBasic, 10)
	data = model.GetUserList()

	code := e.SUCCESS
	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   data,
	}
}
