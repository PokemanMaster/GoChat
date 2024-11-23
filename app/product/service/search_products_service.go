package service

import (
	"IMProject/app/product/model"
	"IMProject/app/product/serializer"
	"IMProject/pkg/e"
	"IMProject/pkg/logging"
	"IMProject/pkg/utils"
	"IMProject/resp"
)

// SearchProductsService 搜索商品的服务
type SearchProductsService struct {
	Search string
}

// Show 搜索商品
func (service *SearchProductsService) Show() resp.Response {

	validSearch, code, err := utils.ValidateSearchInput(service.Search)
	if code != e.SUCCESS {
		logging.Info(err)
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	productParam, code, err := model.SearchProductParam(validSearch)
	if code != e.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProductParams(productParam),
	}
}
