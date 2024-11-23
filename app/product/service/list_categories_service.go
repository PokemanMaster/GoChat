package service

import (
	"IMProject/app/product/model"
	"IMProject/app/product/serializer"
	"IMProject/common/db"
	"IMProject/pkg/e"
	"IMProject/pkg/logging"
	"IMProject/resp"
)

// ListCategoriesService 分类列表服务
type ListCategoriesService struct {
}

// List 商品分类
func (service *ListCategoriesService) List() resp.Response {
	var categories []model.ProductCategory
	code := e.SUCCESS

	if err := db.DB.Find(&categories).Error; err != nil {
		logging.Info(err)
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
