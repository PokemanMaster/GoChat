package api

import (
	service2 "github.com/PokemanMaster/GoChat/server/app/chat/service"
	"github.com/PokemanMaster/GoChat/server/app/chat/ws"
	"github.com/PokemanMaster/GoChat/server/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SendMessage 发送与接收消息
func SendMessage(ctx *gin.Context) {
	ws.Chat(ctx.Writer, ctx.Request)
}

// GetMessage 获取用户A、B的消息
func GetMessage(ctx *gin.Context) {
	services := service2.GetMessageService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.chat.api", "chat.go"))
	} else {
		res := services.Get()
		ctx.JSON(200, res)
	}
}

// Upload 上传文件
func Upload(ctx *gin.Context) {
	services := service2.UploadLocalService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.chat.api", "chat.go"))
	} else {
		res := services.UploadLocal
		ctx.JSON(200, res)
	}
}
