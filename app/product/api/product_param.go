package api

import (
	"IMProject/app/product/service"
	"IMProject/pkg/logging"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

// ListProductsParams 展示商品参数列表
func ListProductsParams(c *gin.Context) {
	services := service.ListProductsParamsService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowProductParam 展示商品详情
func ShowProductParam(ctx *gin.Context) {
	services := service.ShowParamService{}
	res := services.Show(ctx.Param("id"))
	ctx.JSON(200, res)
}
