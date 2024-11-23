package api

import (
	"IMProject/app/product/service"
	"IMProject/pkg/logging"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	services := service.CreateProductService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		logging.Info(err)
	}
}

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

// UpdateProduct 更新商品的接口
func UpdateProduct(c *gin.Context) {
	services := service.UpdateProductService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.Update()
		c.JSON(400, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		logging.Info(err)
	}
}

// DeleteProduct 删除商品的接口
func DeleteProduct(c *gin.Context) {
	services := service.DeleteProductService{}
	res := services.Delete(c.Param("id"))
	c.JSON(200, res)
}

// SearchProducts 搜索商品的接口
func SearchProducts(c *gin.Context) {
	services := service.SearchProductsService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		logging.Info(err)
	}
}
