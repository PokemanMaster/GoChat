package api

import (
	"IMProject/app/group/service"
	"IMProject/pkg/logging"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

// CreateGroup 创建群聊
func CreateGroup(ctx *gin.Context) {
	services := service.CreateGroupService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.Create(ctx)
		ctx.JSON(200, res)
	}
}

// GroupLists 群聊列表
func GroupLists(ctx *gin.Context) {
	services := service.GroupListsService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.List(ctx)
		ctx.JSON(200, res)
	}
}

// JoinGroup 加入群聊
func JoinGroup(ctx *gin.Context) {
	services := service.JoinGroupsService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.Join(ctx)
		ctx.JSON(200, res)
	}
}

// GroupFriendLists 群好友
func GroupFriendLists(ctx *gin.Context) {
	services := service.GroupFriendListsService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.List()
		ctx.JSON(200, res)
	}
}
