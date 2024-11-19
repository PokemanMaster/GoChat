package api

import (
	"IMProject/app/friend/service"
	"IMProject/pkg/logging"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

// CreateFriend 添加用户
func CreateFriend(ctx *gin.Context) {
	services := service.CreateFriendService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.Create()
		ctx.JSON(200, res)
	}
}

// FriendLists 好友列表
func FriendLists(ctx *gin.Context) {
	services := service.FriendListsService{}
	res := services.List(ctx.Param("id"))
	ctx.JSON(200, res)
}

// SearchFriend 搜索好友
func SearchFriend(ctx *gin.Context) {
	services := service.SearchFriendService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.Search()
		ctx.JSON(200, res)
	}
}
