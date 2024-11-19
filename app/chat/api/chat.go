package api

import (
	"IMProject/app/chat/service"
	"IMProject/app/chat/ws"
	"IMProject/pkg/logging"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

// SendMessage 发送与接收消息
func SendMessage(ctx *gin.Context) {
	ws.Chat(ctx.Writer, ctx.Request)
}

// GetMessage 获取用户A、B的消息
func GetMessage(ctx *gin.Context) {
	services := service.GetMessageService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.Get()
		ctx.JSON(200, res)
	}
}

// Upload 上传文件
func Upload(ctx *gin.Context) {
	services := service.UploadLocalService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.UploadLocal
		ctx.JSON(200, res)
	}
}
