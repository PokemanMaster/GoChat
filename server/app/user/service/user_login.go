package service

import (
	"github.com/PokemanMaster/GoChat/server/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/server/common/db"
	"github.com/PokemanMaster/GoChat/server/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/server/pkg/mid"
	"github.com/PokemanMaster/GoChat/server/server/resp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"time"
)

type UserLoginService struct {
	UserName string
	Password string
}

func (service *UserLoginService) UserLogin(ctx *gin.Context) *resp.Response {
	user := model.User{}
	UserName := strings.TrimSpace(service.UserName)
	Password := strings.TrimSpace(service.Password)

	// 检查用户名是否为空
	if UserName == "" {
		zap.L().Info("用户名不能为空", zap.String("app.user.service.user_login", ""))
		return &resp.Response{
			Status: e.ERROR_ACCOUNT_NOT_EMPTY,
			Msg:    e.GetMsg(e.ERROR_ACCOUNT_NOT_EMPTY),
		}
	}

	// 检查密码是否为空
	if Password == "" {
		zap.L().Info("密码不能为空", zap.String("app.user.service.user_login", ""))
		return &resp.Response{
			Status: e.ERROR_PASSWORD_NOT_EMPTY,
			Msg:    e.GetMsg(e.ERROR_PASSWORD_NOT_EMPTY),
		}
	}

	// 根据用户名查找用户
	err := db.DB.Where("user_name = ?", UserName).First(&user).Error
	if err != nil {
		zap.L().Info("账号错误", zap.String("app.user.service.user_login", ""))
		return &resp.Response{
			Status: e.ERROR_ACCOUNT_NOT_EXIST,
			Msg:    e.GetMsg(e.ERROR_ACCOUNT_NOT_EXIST),
		}
	}

	// 验证密码
	if !mid.ValidPassword(Password, user.Salt, user.Password) {
		zap.L().Info("密码错误", zap.String("app.user.service.user_login", ""))
		return &resp.Response{
			Status: e.ERROR_PASSWORD,
			Msg:    e.GetMsg(e.ERROR_PASSWORD),
		}
	}

	// 更新用户登录时间
	user.LoginTime = time.Now()
	if err = db.DB.Save(&user).Error; err != nil {
		zap.L().Error("更新用户失败", zap.String("app.user.service.user_login", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	session := sessions.Default(ctx)
	session.Options(sessions.Options{MaxAge: 3600 * 6})
	session.Set("user_"+UserName, user)
	session.Save()

	// 成功返回用户数据
	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data:   user,
	}
}
