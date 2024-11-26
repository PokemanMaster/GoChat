package api

import (
	service2 "github.com/PokemanMaster/GoChat/server/app/friend/service"
	"github.com/PokemanMaster/GoChat/server/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateFriend 添加用户
func CreateFriend(ctx *gin.Context) {
	services := service2.CreateFriendService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.friend.api.friend.go", err.Error()))
	} else {
		res := services.Create()
		ctx.JSON(200, res)
	}
}

// FriendLists 好友列表
func FriendLists(ctx *gin.Context) {
	services := service2.FriendListsService{}
	res := services.List(ctx.Param("id"))
	ctx.JSON(200, res)
}

// SearchFriend 搜索好友
func SearchFriend(ctx *gin.Context) {
	services := service2.SearchFriendService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.friend.api.friend.go", err.Error()))
	} else {
		res := services.Search()
		ctx.JSON(200, res)
	}
}

// DeleteFriend 删除好友
func DeleteFriend(ctx *gin.Context) {
	services := service2.DeleteFriendService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.friend.api.friend.go", err.Error()))
	} else {
		res := services.Delete()
		ctx.JSON(200, res)
	}
}
