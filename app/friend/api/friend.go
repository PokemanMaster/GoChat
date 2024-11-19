package api

import (
	service2 "IMProject/app/friend/service"
	"IMProject/app/user/service"
	"IMProject/pkg/logging"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

// CreateFriend 添加用户
func CreateFriend(ctx *gin.Context) {
	services := service2.CreateFriendService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.Create(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// FriendLists 好友列表
func FriendLists(ctx *gin.Context) {
	services := service2.FriendListsService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.List(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// SearchFriend 搜索好友
func SearchFriend(ctx *gin.Context) {
	services := service.SearchFriendService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.Search(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}
