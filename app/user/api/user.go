package api

import (
	"github.com/PokemanMaster/GoChat/app/user/service"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/PokemanMaster/GoChat/resp"
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
		res := services.UserLogin(ctx)
		ctx.JSON(200, res)
	}
}

// UserLogout 用户登出接口
func UserLogout(ctx *gin.Context) {
	services := service.UserLogoutService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.UserLogout(ctx)
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

// CaptchaImage 给用户返回base64码的图片
func CaptchaImage(ctx *gin.Context) {
	services := service.UserCategoryService{}
	res := services.UserCategoryImages()
	ctx.JSON(200, res)
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
