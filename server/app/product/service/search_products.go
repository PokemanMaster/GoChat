package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/product/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/mid"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
)

// SearchProductsService 搜索商品的服务
type SearchProductsService struct {
	Search string
}

// Show 搜索商品
func (service *SearchProductsService) Show() resp.Response {
	var products []model.Product
	var productParams []model.ProductParam
	var result []serializer.ProductParamSerialization

	// 验证搜索输入
	validSearch, err := mid.ValidateSearchInput(service.Search)
	if err != nil {
		zap.L().Error("服务被攻击了", zap.String("app.product.service", err.Error()))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	// 搜索商品
	err = db.DB.Where("name LIKE ?", "%"+validSearch+"%").Find(&products).Error
	if err != nil {
		zap.L().Error("搜索失败", zap.String("app.product.service", err.Error()))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	// 遍历商品，查找每个商品对应的最小价格
	for _, product := range products {
		err = db.DB.Where("product_id = ?", product.ID).Order("price asc").Find(&productParams).Error
		if err != nil {
			zap.L().Error("查询商品参数错误", zap.String("app.product.service", err.Error()))
			continue
		}
		if len(productParams) > 0 {
			cheapestParam := productParams[0] // 已按价格升序排序，价格最小的在第一个
			result = append(result, serializer.ProductParamSerialization{
				ID:        cheapestParam.ID,
				ProductID: cheapestParam.ProductID,
				Name:      product.Name,
				Price:     cheapestParam.Price,
				Image:     product.Image,
				Saleable:  product.Saleable,
				SoldCount: cheapestParam.SoldCount,
			})
		}
	}

	// 返回搜索结果
	return resp.BuildResponseTotal(result, uint(len(result)))
}
