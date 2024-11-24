package service

import (
	"github.com/PokemanMaster/GoChat/app/user/model"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
	"github.com/gin-gonic/gin"
)

type UserListsService struct {
}

func (service *UserListsService) List(c *gin.Context) *resp.Response {
	data := make([]*model.User, 10)
	data = model.GetUserList()

	code := e.SUCCESS
	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   data,
	}
}
