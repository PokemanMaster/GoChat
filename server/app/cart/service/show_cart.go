package service

import (
	"context"
	"encoding/json"
	"github.com/PokemanMaster/GoChat/v1/server/app/cart/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/cart/serializer"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"

	"time"
)

type ShowCartService struct{}

// Show 获取购物车的商品
func (service *ShowCartService) Show(ctx context.Context, id string) resp.Response {
	CartRedisKey := "ShowCart_" + id
	var carts []model.Cart

	// 查询 redis
	CartsCache, err := cache.RC.Get(ctx, CartRedisKey).Result()
	if err == nil && CartsCache != "" {
		if err := json.Unmarshal([]byte(CartsCache), &carts); err != nil {
			zap.L().Error("Cart 缓存数据解析失败", zap.String("app.cart.service", "show_cart.go"))
			return resp.Response{
				Status: e.ERROR_UNMARSHAL_JSON,
				Msg:    e.GetMsg(e.ERROR_UNMARSHAL_JSON),
				Error:  err.Error(),
			}
		}
		resp.BuildResponseTotal(serializer.BuildCarts(carts), uint(len(carts)))
	}

	// 如果缓存未命中，则从数据库查询
	CartsData, code := model.ListCart(id)
	if code != e.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 将数据库查询到的购物车数据存入 Redis
	CartsJSON, _ := json.Marshal(CartsData)
	err = cache.RC.Set(ctx, CartRedisKey, CartsJSON, 24*time.Hour).Err()
	if err != nil {
		zap.L().Error("Cart 缓存创建/更新失败", zap.String("app.cart.service", "show_cart.go"))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	return resp.BuildResponseTotal(serializer.BuildCarts(carts), uint(len(carts)))
}
