package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/category/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
)

// ListCategoriesService 分类列表服务
type ListCategoriesService struct {
}

// List 商品分类
func (service *ListCategoriesService) List() resp.Response {
	var categories []model.ProductCategory
	code := e.SUCCESS

	if err := db.DB.Find(&categories).Error; err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProductCategorys(categories),
	}
}
