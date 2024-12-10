package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/PokemanMaster/GoChat/v1/server/app/product/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache" // 引入 Redis 缓存包
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
)

// ListRankingService 展示排行的服务
type ListRankingService struct {
}

// List 获取排行
func (service *ListRankingService) List(ctx context.Context) resp.Response {
	var products []model.Product

	cacheKey := "product_rankings" // 缓存的键

	// 尝试从 Redis 获取缓存数据
	cachedData, err := cache.RC.Get(ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		// 缓存存在，解码并返回缓存结果
		if err := json.Unmarshal([]byte(cachedData), &products); err != nil {
			zap.L().Error("解码缓存数据失败", zap.String("app.product.ranking", err.Error()))
			return resp.Response{
				Status: e.ERROR_DATABASE,
				Msg:    e.GetMsg(e.ERROR_DATABASE),
				Error:  err.Error(),
			}
		}

		// 缓存命中，返回结果
		return resp.Response{
			Status: e.SUCCESS,
			Msg:    e.GetMsg(e.SUCCESS),
			Data:   products,
		}
	}

	// 如果缓存不存在或出错，则从数据库中查询
	err = db.DB.Order("rating DESC").Limit(10).Find(&products).Error
	if err != nil {
		zap.L().Error("查看排行数据错误", zap.String("app.product.ranking", err.Error()))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	// 查询成功后，将数据缓存到 Redis，设置 TTL 为一天（86400秒）
	cachedDataBytes, err := json.Marshal(products)
	if err != nil {
		zap.L().Error("序列化数据失败", zap.String("app.product.ranking", err.Error()))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	err = cache.RC.Set(ctx, cacheKey, cachedDataBytes, 24*time.Hour).Err() // 设置缓存，过期时间为24小时
	if err != nil {
		zap.L().Error("缓存写入失败", zap.String("app.product.ranking", err.Error()))
	}

	// 返回数据库查询的结果
	return resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data:   products,
	}
}
