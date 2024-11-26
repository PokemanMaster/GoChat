package service

import (
	"context"
	"encoding/json"
	"github.com/PokemanMaster/GoChat/server/app/cart/model"
	"github.com/PokemanMaster/GoChat/server/app/cart/serializer"
	"github.com/PokemanMaster/GoChat/server/common/cache"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
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
				Status: e2.ERROR_UNMARSHAL_JSON,
				Msg:    e2.GetMsg(e2.ERROR_UNMARSHAL_JSON),
				Error:  err.Error(),
			}
		}
		resp.BuildResponseTotal(serializer.BuildCarts(carts), uint(len(carts)))
	}

	// 如果缓存未命中，则从数据库查询
	CartsData, code := model.ListCart(id)
	if code != e2.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}

	// 将数据库查询到的购物车数据存入 Redis
	CartsJSON, _ := json.Marshal(CartsData)
	err = cache.RC.Set(ctx, CartRedisKey, CartsJSON, 24*time.Hour).Err()
	if err != nil {
		zap.L().Error("Cart 缓存创建/更新失败", zap.String("app.cart.service", "show_cart.go"))
		return resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	return resp.BuildResponseTotal(serializer.BuildCarts(carts), uint(len(carts)))
}
