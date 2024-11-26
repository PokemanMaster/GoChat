package api

import (
	"github.com/PokemanMaster/GoChat/server/server/app/category/service"
	"github.com/PokemanMaster/GoChat/server/server/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ListCategories 商品分类列表接口
func ListCategories(c *gin.Context) {
	services := service.ListCategoriesService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.chat.api", "chat.go"))
	}
}
