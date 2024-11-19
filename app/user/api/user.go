package api

import (
	"IMProject/app/user/service"
	"IMProject/pkg/logging"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(ctx *gin.Context) {
	services := service.UserRegisterService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.UserRegister()
		ctx.JSON(200, res)
	}
}

// UserLogin 用户登录接口
func UserLogin(ctx *gin.Context) {
	services := service.UserLoginService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.UserLogin()
		ctx.JSON(200, res)
	}
}

// UserUpdate 修改用户信息
func UserUpdate(ctx *gin.Context) {
	services := service.UserUpdateService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.UserUpdate(ctx)
		ctx.JSON(200, res)
	}
}

// UserInfo 获取用户信息
func UserInfo(ctx *gin.Context) {
	services := service.UserInfoService{}
	res := services.UserInfo(ctx.Param("id"))
	ctx.JSON(200, res)
}

// UserLists 用户列表
func UserLists(ctx *gin.Context) {
	services := service.UserListsService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.List(ctx)
		ctx.JSON(200, res)
	}
}
