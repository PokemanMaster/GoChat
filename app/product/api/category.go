package api

import (
	"IMProject/app/product/service"
	"IMProject/pkg/logging"
	"IMProject/resp"
	"github.com/gin-gonic/gin"
)

// ListCategories 商品分类列表接口
func ListCategories(c *gin.Context) {
	services := service.ListCategoriesService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
		logging.Info(err)
	}
}
