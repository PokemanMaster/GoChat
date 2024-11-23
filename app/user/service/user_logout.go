package service

import (
	"IMProject/pkg/e"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

type UserLogoutService struct {
	token string
}

func (service *UserLogoutService) UserLogout(ctx *gin.Context) *resp.Response {
	code := 200

	ctx.SetCookie("MyCookie", service.token, -1, "/", "localhost", false, true)
	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
