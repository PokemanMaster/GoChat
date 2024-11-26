package api

import (
	"github.com/PokemanMaster/GoChat/server/server/app/payment/service"
	"github.com/PokemanMaster/GoChat/server/server/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePay 创建支付
func CreatePay(ctx *gin.Context) {
	services := service.CreatePayService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.chat.api", "chat.go"))
	} else {
		res := services.Create(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}
