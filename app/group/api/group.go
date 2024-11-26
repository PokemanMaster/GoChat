package api

import (
	service2 "github.com/PokemanMaster/GoChat/server/app/group/service"
	"github.com/PokemanMaster/GoChat/server/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateGroup 创建群聊
func CreateGroup(ctx *gin.Context) {
	services := service2.CreateGroupService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.group.api", "group.go"))
	} else {
		res := services.Create()
		ctx.JSON(200, res)
	}
}

// GroupLists 群聊列表
func GroupLists(ctx *gin.Context) {
	services := service2.GroupListsService{}
	res := services.List(ctx.Param("id"))
	ctx.JSON(200, res)
}

// JoinGroup 加入群聊
func JoinGroup(ctx *gin.Context) {
	services := service2.JoinGroupService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.group.api", "group.go"))
	} else {
		res := services.Join()
		ctx.JSON(200, res)
	}
}

// GroupFriendLists 群好友
func GroupFriendLists(ctx *gin.Context) {
	services := service2.GroupFriendListsService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.group.api", "group.go"))
	} else {
		res := services.List()
		ctx.JSON(200, res)
	}
}
