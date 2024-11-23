package api

import (
	"IMProject/app/product/service"
	"IMProject/pkg/logging"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

// CreateCarousel 创建轮播图
func CreateCarousel(c *gin.Context) {
	services := service.CreateCarouselService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		logging.Info(err)
	}
}

// ListCarousels 轮播图列表接口
func ListCarousels(c *gin.Context) {
	services := service.ListCarouselsService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		logging.Info(err)
	}
}
