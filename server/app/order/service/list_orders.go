package service

import (
	"context"
	"encoding/json"
	"github.com/PokemanMaster/GoChat/v1/server/app/order/build"
	"github.com/PokemanMaster/GoChat/v1/server/app/order/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"

	"time"
)

// ListOrdersService 订单详情的服务
type ListOrdersService struct {
}

func (service *ListOrdersService) List(ctx context.Context, id string) resp.Response {
	OrdersRedisKey := "ShowOrder_" + id
	var orders []model.Order

	// 查询缓存数据
	ordersCache, err := cache.RC.Get(ctx, OrdersRedisKey).Result()
	if err == nil && ordersCache != "" {
		if err = json.Unmarshal([]byte(ordersCache), &orders); err != nil {
			zap.L().Error("查询缓存数据失败", zap.String("app.order.service.order", err.Error()))
			return resp.Response{
				Status: e.ERROR_UNMARSHAL_JSON,
				Msg:    e.GetMsg(e.ERROR_UNMARSHAL_JSON),
				Error:  err.Error(),
			}
		}
		resp.BuildResponseTotal(build.ResUserOrders(orders), uint(len(orders)))
	}

	// 如果缓存未命中，则从数据库查询
	err = db.DB.Where("user_id=?", id).Find(&orders).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	// 将数据库查询到的购物车数据存入 Redis
	OrdersJSON, _ := json.Marshal(orders)
	err = cache.RC.Set(ctx, OrdersRedisKey, OrdersJSON, 24*time.Hour).Err()
	if err != nil {
		zap.L().Error("存储缓存数据失败", zap.String("app.order.service.order", err.Error()))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	return resp.BuildResponseTotal(build.ResUserOrders(orders), uint(len(orders)))
}
