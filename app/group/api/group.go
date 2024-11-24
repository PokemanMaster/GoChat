package api

import (
	"github.com/PokemanMaster/GoChat/app/group/service"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/PokemanMaster/GoChat/resp"
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
		res := services.Create()
		ctx.JSON(200, res)
	}
}

// GroupLists 群聊列表
func GroupLists(ctx *gin.Context) {
	services := service.GroupListsService{}
	res := services.List(ctx.Param("id"))
	ctx.JSON(200, res)
}

// JoinGroup 加入群聊
func JoinGroup(ctx *gin.Context) {
	services := service.JoinGroupService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.Join()
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
