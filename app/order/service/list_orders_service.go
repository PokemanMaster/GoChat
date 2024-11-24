package service

import (
	"context"
	"encoding/json"
	"github.com/PokemanMaster/GoChat/app/order/build"
	"github.com/PokemanMaster/GoChat/app/order/model"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/PokemanMaster/GoChat/resp"

	"time"
)

// ListOrdersService 订单详情的服务
type ListOrdersService struct {
}

func (service *ListOrdersService) List(ctx context.Context, id string) resp.Response {
	OrdersRedisKey := "ShowOrder_" + id
	var orders []model.Order

	// 查询 redis
	OrdersCache, err := cache.RC.Get(ctx, OrdersRedisKey).Result()
	if err == nil && OrdersCache != "" {
		if err := json.Unmarshal([]byte(OrdersCache), &orders); err != nil {
			logging.Info("订单JSON解析失败", err)
			return resp.Response{
				Status: e.ERROR_UNMARSHAL_JSON,
				Msg:    e.GetMsg(e.ERROR_UNMARSHAL_JSON),
				Error:  err.Error(),
			}
		}
		resp.BuildListResponse(build.ResUserOrders(orders), uint(len(orders)))
	}

	// 如果缓存未命中，则从数据库查询
	OrdersData, code := model.ListOrder(id)
	if code != e.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 将数据库查询到的购物车数据存入 Redis
	OrdersJSON, _ := json.Marshal(OrdersData)
	err = cache.RC.Set(ctx, OrdersRedisKey, OrdersJSON, 24*time.Hour).Err()
	if err != nil {
		logging.Info("Order 缓存创建/更新失败", err)
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(code),
		}
	}

	return resp.BuildListResponse(build.ResUserOrders(OrdersData), uint(len(orders)))
}
