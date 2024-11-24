package service

import (
	"github.com/PokemanMaster/GoChat/app/product/model"
	"github.com/PokemanMaster/GoChat/app/product/serializer"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/PokemanMaster/GoChat/resp"
)

// ListProductsService 视频列表服务
type ListProductsService struct {
	Limit      int // 限制展示商品的个数
	Start      int // 选择开始的序号
	CategoryID uint
}

// List 各个商品列表
func (service *ListProductsService) List() resp.Response {
	var products []model.Product
	var total int64
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}

	// 0：推荐、1：食品、2：水果、3：男装、4：电脑、5:医药
	if service.CategoryID == 0 {
		if err := db.DB.Model(model.ProductParam{}).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := db.DB.Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := db.DB.Model(model.Product{}).Where("category_id=?", service.CategoryID).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := db.DB.Where("category_id=?", service.CategoryID).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return resp.BuildListResponse(serializer.BuildProducts(products), uint(total))
}
