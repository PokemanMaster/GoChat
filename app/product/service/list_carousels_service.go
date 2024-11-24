package service

import (
	"context"
	"encoding/json"
	"github.com/PokemanMaster/GoChat/app/product/model"
	"github.com/PokemanMaster/GoChat/app/product/serializer"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
	"go.uber.org/zap"

	"time"
)

type ListCarouselsService struct {
}

// List 视频列表
func (service *ListCarouselsService) List(ctx context.Context) resp.Response {
	var carousels []model.Carousel
	code := e.SUCCESS

	// Redis key
	redisKey := "ListCarousels"

	// 尝试从 Redis 中获取数据
	cachedData, err := cache.RC.Get(ctx, redisKey).Result()

	if err == nil && cachedData != "" {
		// 如果缓存命中，直接反序列化并返回
		err = json.Unmarshal([]byte(cachedData), &carousels)
		if err != nil {
			zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
			code = e.ERROR
			return resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		// 返回缓存的数据
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   serializer.BuildCarousels(carousels),
		}
	}

	// 如果 Redis 中没有数据，查询数据库
	if err := db.DB.Find(&carousels).Error; err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 将查询结果缓存到 Redis，设置过期时间
	cachedDataBytes, err := json.Marshal(carousels)
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e.ERROR
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 将数据写入 Redis，设置 TTL 为 1 小时
	err = cache.RC.Set(ctx, redisKey, cachedDataBytes, time.Hour*114514).Err()
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e.ERROR
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 返回数据库查询的数据
	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarousels(carousels),
	}
}
