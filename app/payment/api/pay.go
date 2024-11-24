package api

import (
	"github.com/PokemanMaster/GoChat/app/payment/service"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/PokemanMaster/GoChat/resp"
	"github.com/gin-gonic/gin"
)

// CreatePay 创建支付
func CreatePay(ctx *gin.Context) {
	services := service.CreatePayService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.Create(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}
