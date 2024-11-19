package service

import (
	"IMProject/app/user/model"
	"IMProject/pkg/e"
	"IMProject/pkg/utils"
	"IMProject/resp"
	"strings"
)

type UserLoginService struct {
	UserName string
	Password string
}

func (service *UserLoginService) UserLogin() *resp.Response {
	user := model.UserBasic{}
	user.Name = strings.TrimSpace(service.UserName) // 去除多余的空格
	password := service.Password

	// 检查用户名和密码是否为空
	if user.Name == "" {
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
	user = model.FindUserByName(user.Name)
	if user.Name == "" {
		return &resp.Response{
			Status: e.ERROR_ACCOUNT_NOT_EXIST,
			Msg:    e.GetMsg(e.ERROR_ACCOUNT_NOT_EXIST),
		}
	}

	// 验证密码
	if !utils.ValidPassword(password, user.Salt, user.PassWord) {
		return &resp.Response{
			Status: e.ERROR_PASSWORD,
			Msg:    e.GetMsg(e.ERROR_PASSWORD),
		}
	}

	// 生成加密密码并再次查询确认
	encryptedPwd := utils.MakePassword(password, user.Salt)
	user = model.FindUserByNameAndPwd(user.Name, encryptedPwd)
	if user.Name == "" {
		return &resp.Response{
			Status: e.ERROR_PASSWORD,
			Msg:    e.GetMsg(e.ERROR_PASSWORD),
		}
	}

	// 成功返回用户数据
	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data:   user,
	}
}
