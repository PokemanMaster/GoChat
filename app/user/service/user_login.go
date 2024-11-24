package service

import (
	"github.com/PokemanMaster/GoChat/app/user/model"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/mid"
	"github.com/PokemanMaster/GoChat/resp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserLoginService struct {
	UserName string
	Password string
}

func (service *UserLoginService) UserLogin(ctx *gin.Context) *resp.Response {

	user := model.User{}
	user.UserName = strings.TrimSpace(service.UserName)
	password := service.Password

	// 检查用户名和密码是否为空
	if user.UserName == "" {
		return &resp.Response{
			Status: e.ERROR_ACCOUNT_NOT_EMPTY,
			Msg:    e.GetMsg(e.ERROR_ACCOUNT_NOT_EMPTY),
		}
	}

	// 检查用户名和密码是否为空
	if password == "" {
		return &resp.Response{
			Status: e.ERROR_PASSWORD_NOT_EMPTY,
			Msg:    e.GetMsg(e.ERROR_PASSWORD_NOT_EMPTY),
		}
	}

	// 根据用户名查找用户
	user = model.FindUserByName(user.UserName)
	if user.UserName == "" {
		return &resp.Response{
			Status: e.ERROR_ACCOUNT_NOT_EXIST,
			Msg:    e.GetMsg(e.ERROR_ACCOUNT_NOT_EXIST),
		}
	}

	// 验证密码
	if !mid.ValidPassword(password, user.Salt, user.Password) {
		return &resp.Response{
			Status: e.ERROR_PASSWORD,
			Msg:    e.GetMsg(e.ERROR_PASSWORD),
		}
	}

	// 生成加密密码并再次查询确认
	encryptedPwd := mid.MakePassword(password, user.Salt)
	user = model.FindUserByNameAndPwd(user.UserName, encryptedPwd)
	if user.UserName == "" {
		return &resp.Response{
			Status: e.ERROR_PASSWORD,
			Msg:    e.GetMsg(e.ERROR_PASSWORD),
		}
	}

	session := sessions.Default(ctx)
	session.Options(sessions.Options{MaxAge: 3600 * 6})
	session.Set("user_"+service.UserName, user)
	session.Save()

	// 成功返回用户数据
	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data:   user,
	}
}
