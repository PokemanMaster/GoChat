package service

import (
	"errors"
	"fmt"
	"github.com/PokemanMaster/GoChat/server/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/server/common/db"
	"github.com/PokemanMaster/GoChat/server/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/server/pkg/mid"
	"github.com/PokemanMaster/GoChat/server/server/resp"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/rand"
	"strings"
	"time"
)

type UserRegisterService struct {
	UserName   string
	Password   string
	RePassword string
}

func (service *UserRegisterService) UserRegister() *resp.Response {
	user := model.User{}
	UserName := strings.TrimSpace(service.UserName)
	Password := strings.TrimSpace(service.Password)
	RePassword := strings.TrimSpace(service.RePassword)

	// 检查用户名是否为空
	if UserName == "" {
		zap.L().Info("用户名不能为空", zap.String("app.user.service.user_register", ""))
		return &resp.Response{
			Status: e.ERROR_ACCOUNT_NOT_EMPTY,
			Msg:    e.GetMsg(e.ERROR_ACCOUNT_NOT_EMPTY),
		}
	}

	// 检查密码和重复密码是否为空
	if Password == "" || RePassword == "" {
		zap.L().Info("密码不能为空", zap.String("app.user.service.user_register", ""))
		return &resp.Response{
			Status: e.ERROR_PASSWORD_NOT_EMPTY,
			Msg:    e.GetMsg(e.ERROR_PASSWORD_NOT_EMPTY),
		}
	}

	// 检查密码和确认密码是否一致
	if Password != RePassword {
		zap.L().Info("两次密码不一致", zap.String("app.user.service.user_register", ""))
		return &resp.Response{
			Status: e.ERROR_PASSWORD_CONFIRM,
			Msg:    e.GetMsg(e.ERROR_PASSWORD_CONFIRM),
		}
	}

	// 查询用户是否存在
	err := db.DB.Where("user_name = ?", UserName).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Info("用户名可用", zap.String("app.user.service.user_register", ""))
		} else {
			zap.L().Error("查询用户失败", zap.String("app.user.service.user_register", ""))
			return &resp.Response{
				Status: e.ERROR_DATABASE,
				Msg:    e.GetMsg(e.ERROR_DATABASE),
				Error:  err.Error(),
			}
		}
	} else {
		zap.L().Info("用户名已存在", zap.String("app.user.service.user_register", ""))
		return &resp.Response{
			Status: e.ERROR_ACCOUNT_EXIST,
			Msg:    e.GetMsg(e.ERROR_ACCOUNT_EXIST),
		}
	}

	// 加密密码 并 创建用户
	Salt := fmt.Sprintf("%06d", rand.Int31())
	user.UserName = UserName
	user.Password = mid.MakePassword(Password, Salt)
	user.Salt = Salt
	user.LevelID = 1
	user.Money = 0
	user.HeartbeatTime = time.Now()
	err = db.DB.Create(&user).Error
	if err != nil {
		zap.L().Error("创建用户失败", zap.String("app.user.service.user_register", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	// 注册成功
	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
