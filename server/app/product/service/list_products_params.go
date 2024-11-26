package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PokemanMaster/GoChat/v1/server/app/product/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
	"time"
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
	var productsParam []model.ProductParam
	var total int64
	code := e.SUCCESS

	// 默认展示数量
	if service.Limit == 0 {
		service.Limit = 15
	}

	// 生成 Redis 缓存的键
	redisKey := generateRedisKey(service.CategoryID, service.Start, service.Limit)

	// 尝试从 Redis 中获取缓存的数据
	cachedData, err := cache.RC.Get(ctx, redisKey).Result()
	if err == nil && cachedData != "" {
		// 如果缓存命中，直接返回数据
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		var cachedProducts []model.ProductParam
		if err := json.Unmarshal([]byte(cachedData), &cachedProducts); err == nil {
			return resp.BuildResponseTotal(serializer.BuildProductParams(cachedProducts), uint(len(cachedProducts)))
		} else {
			zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		}
	} else {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
	}

	// 如果分类ID为 0，查找所有商品
	if service.CategoryID == 0 {
		if err := db.DB.Limit(service.Limit).Offset(service.Start).Find(&productsParam).Error; err != nil {
			zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
			code = e.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else { // 查找对应分类的商品
		var productIDs []uint
		if err := db.DB.Model(&model.Product{}).Where("category_id = ?", service.CategoryID).Pluck("id", &productIDs).Error; err != nil {
			zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
			code = e.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		// 根据 productIDs 查找 ProductParam
		if err := db.DB.Model(&model.ProductParam{}).Where("product_id IN (?)", productIDs).Count(&total).Error; err != nil {
			zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
			code = e.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := db.DB.Where("product_id IN (?)", productIDs).Limit(service.Limit).Offset(service.Start).Find(&productsParam).Error; err != nil {
			zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
			code = e.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	// 将查询结果缓存到 Redis 中，设置过期时间为 10 分钟
	productsJSON, err := json.Marshal(productsParam)
	if err == nil {
		err = cache.RC.Set(ctx, redisKey, productsJSON, 10*time.Minute).Err()
		if err != nil {
			zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		}
	} else {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
	}

	return resp.BuildResponseTotal(serializer.BuildProductParams(productsParam), uint(len(productsParam)))
}
