package api

import (
	"github.com/PokemanMaster/GoChat/app/product/service"
	"github.com/PokemanMaster/GoChat/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateCarousel 创建轮播图
func CreateCarousel(c *gin.Context) {
	services := service.CreateCarouselService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.chat.api", "chat.go"))
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
		zap.L().Error("请求参数错误", zap.String("app.chat.api", "chat.go"))
	}
}
