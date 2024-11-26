package service

import (
	model2 "github.com/PokemanMaster/GoChat/server/app/product/model"
	"github.com/PokemanMaster/GoChat/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
	"go.uber.org/zap"
)

// ListProductsService 视频列表服务
type ListProductsService struct {
	Limit      int // 限制展示商品的个数
	Start      int // 选择开始的序号
	CategoryID uint
}

// List 各个商品列表
func (service *ListProductsService) List() resp.Response {
	var products []model2.Product
	var total int64
	code := e2.SUCCESS

	if service.Limit <= 0 {
		service.Limit = 15
	}

	// 0：推荐、1：食品、2：水果、3：男装、4：电脑、5:医药
	if service.CategoryID == 0 {
		if err := db.DB.Model(model2.ProductParam{}).Count(&total).Error; err != nil {
			zap.L().Error("请求失败", zap.String("app.product.service.list_products", err.Error()))
			return resp.Response{
				Status: e2.ERROR_DATABASE,
				Msg:    e2.GetMsg(e2.ERROR_DATABASE),
				Error:  err.Error(),
			}
		}

		if err := db.DB.Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			zap.L().Error("请求失败", zap.String("app.product.service.list_products", err.Error()))
			code = e2.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e2.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := db.DB.Model(model2.Product{}).Where("category_id=?", service.CategoryID).Count(&total).Error; err != nil {
			zap.L().Error("请求失败", zap.String("app.product.service.list_products", err.Error()))
			code = e2.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e2.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := db.DB.Where("category_id=?", service.CategoryID).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			zap.L().Error("请求失败", zap.String("app.product.service.list_products", err.Error()))
			code = e2.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e2.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return resp.BuildResponseTotal(serializer.BuildProducts(products), uint(total))
}
