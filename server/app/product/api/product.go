package api

import (
	"github.com/PokemanMaster/GoChat/server/server/app/product/service"
	"github.com/PokemanMaster/GoChat/server/server/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ListProducts 展示商品详情列表
func ListProducts(c *gin.Context) {
	services := service.ListProductsService{}
	res := services.List()
	c.JSON(200, res)
}

// ShowProduct 展示商品详情
func ShowProduct(ctx *gin.Context) {
	services := service.ShowProductService{}
	res := services.Show(ctx, ctx.Param("id"))
	ctx.JSON(200, res)
}

// SearchProducts 搜索商品的接口
func SearchProducts(c *gin.Context) {
	services := service.SearchProductsService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.chat.api", "chat.go"))
	}
}
