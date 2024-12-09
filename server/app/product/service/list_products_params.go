package service

import (
	"context"
	"fmt"
	"github.com/PokemanMaster/GoChat/v1/server/app/product/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
)

// ListProductsParamsService 商品列表服务
type ListProductsParamsService struct {
	Limit      int  // 限制展示商品的个数
	Start      int  // 选择开始的序号
	CategoryID uint // 商品分类id
}

// 生成 Redis 键
func generateRedisKey(categoryID uint, start, limit int) string {
	return fmt.Sprintf("ProductList_Category_%d_Start_%d_Limit_%d", categoryID, start, limit)
}

// List 各个商品列表
func (service *ListProductsParamsService) List(ctx context.Context) resp.Response {
	var products []model.Product
	var productParams []model.ProductParam
	var result []serializer.ProductParamSerialization

	// 默认展示数量
	if service.Limit == 0 {
		service.Limit = 10
	}

	// 生成 Redis 缓存的键
	//redisKey := generateRedisKey(service.CategoryID, service.Start, service.Limit)

	// 获取缓存的数据
	//cacheData, err := cache.RC.Get(ctx, redisKey).Result()
	//if err == nil && cacheData != "" {
	//	err = json.Unmarshal([]byte(cacheData), &result)
	//	if err != nil {
	//		zap.L().Error("JSON解析失败", zap.String("app.product.service", err.Error()))
	//	} else {
	//		return resp.BuildResponseTotal(result, uint(len(result)))
	//	}
	//}

	// 先根据商品分类查找对应的商品信息
	err := db.DB.Where("category_id = ?", service.CategoryID).Limit(service.Limit).Offset(service.Start).Find(&products).Error
	if err != nil {
		zap.L().Error("查询商品错误", zap.String("app.product.service", err.Error()))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	// 对每个商品查询对应的 ProductParam 信息，并取金钱最少的一个
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

	// 存储缓存的数据
	//productsJSON, err := json.Marshal(result)
	//if err != nil {
	//	zap.L().Error("序列化商品列表错误", zap.String("app.product.service", err.Error()))
	//} else {
	//	err = cache.RC.Set(ctx, redisKey, productsJSON, 10*time.Minute).Err()
	//	if err != nil {
	//		zap.L().Error("缓存商品列表错误", zap.String("app.product.service", err.Error()))
	//	}
	//}

	return resp.BuildResponseTotal(result, uint(len(result)))
}
