package service

import (
	"fmt"
	"github.com/PokemanMaster/GoChat/app/user/model"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/mid"
	"github.com/PokemanMaster/GoChat/resp"
	"math/rand"
	"strings"
	"time"
)

type UserRegisterService struct {
	UserName string
	Password string
	Identity string
}

func (service *UserRegisterService) UserRegister() *resp.Response {
	user := model.UserBasic{}
	user.UserName = strings.TrimSpace(service.UserName) // 去除多余的空格
	password := service.Password
	repassword := service.Identity

	// 检查用户名和密码是否为空
	if user.UserName == "" || password == "" || repassword == "" {
		return &resp.Response{
			Status: e.ERROR_ACCOUNT_NOT_EXIST,
			Msg:    e.GetMsg(e.ERROR_ACCOUNT_NOT_EXIST),
		}
	}

	if password == "" || repassword == "" {
		return &resp.Response{
			Status: e.ERROR_PASSWORD_NOT_EMPTY,
			Msg:    e.GetMsg(e.ERROR_PASSWORD_NOT_EMPTY),
		}
	}

	// 通过 ORM 查找用户，防止 SQL 注入
	data := model.FindUserByName(user.UserName)
	if data.UserName != "" {
		return &resp.Response{
			Status: e.ERROR_ACCOUNT_NOT_EXIST,
			Msg:    e.GetMsg(e.ERROR_ACCOUNT_NOT_EXIST),
		}
	}

	// 检查密码和确认密码是否一致
	if password != repassword {
		code := e.ERROR_PASSWORD_CONFIRM
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 加密密码
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Password = mid.MakePassword(password, salt)
	user.Salt = salt
	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()

	// 创建用户
	model.CreateUser(user)

	// 注册成功
	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
