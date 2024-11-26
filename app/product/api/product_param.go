package api

import (
	service2 "github.com/PokemanMaster/GoChat/server/app/product/service"
	"github.com/PokemanMaster/GoChat/server/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ListProductsParams 展示商品参数列表
func ListProductsParams(c *gin.Context) {
	services := service2.ListProductsParamsService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.chat.api", "chat.go"))
	}
}

// ShowProductParam 展示商品详情
func ShowProductParam(ctx *gin.Context) {
	services := service2.ShowParamService{}
	res := services.Show(ctx.Param("id"))
	ctx.JSON(200, res)
}
