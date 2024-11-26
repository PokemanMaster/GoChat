package service

import (
	"context"
	"github.com/PokemanMaster/GoChat/server/server/app/product/model"
	"github.com/PokemanMaster/GoChat/server/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/server/server/common/db"
	"github.com/PokemanMaster/GoChat/server/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/server/resp"
	"go.uber.org/zap"
)

// 商品类

// ShowProductService 商品详情的服务
type ShowProductService struct {
}

// Show 商品
func (service *ShowProductService) Show(ctx context.Context, id string) resp.Response {
	var product model.Product
	code := e.SUCCESS
	err := db.DB.First(&product, id).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//增加点击数
	product.AddView(ctx)
	if product.CategoryID == 2 || product.CategoryID == 3 {
		product.AddElecRank(ctx)
	}
	if product.CategoryID == 5 || product.CategoryID == 6 || product.CategoryID == 7 || product.CategoryID == 8 {
		product.AddAcceRank(ctx)
	}

	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
