package service

import (
	"github.com/PokemanMaster/GoChat/server/app/category/model"
	"github.com/PokemanMaster/GoChat/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
	"go.uber.org/zap"
)

// ListCategoriesService 分类列表服务
type ListCategoriesService struct {
}

// List 商品分类
func (service *ListCategoriesService) List() resp.Response {
	var categories []model.ProductCategory
	code := e2.SUCCESS

	if err := db.DB.Find(&categories).Error; err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e2.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return resp.Response{
		Status: code,
		Msg:    e2.GetMsg(code),
		Data:   serializer.BuildProductCategorys(categories),
	}
}
