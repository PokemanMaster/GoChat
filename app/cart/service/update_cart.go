package service

import (
	"context"
	"encoding/json"
	"github.com/PokemanMaster/GoChat/app/cart/model"
	"github.com/PokemanMaster/GoChat/app/cart/serializer"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
	"go.uber.org/zap"

	"strconv"
	"time"
)

// UpdateCartService 购物车修改的服务
type UpdateCartService struct {
	UserID    uint
	ProductID uint
	Num       uint
}

// Update 修改购物车信息
func (service *UpdateCartService) Update(ctx context.Context) *resp.Response {
	// 定义 Redis 的 key
	cartRedisKey := "ShowCart_" + strconv.Itoa(int(service.UserID))

	// 从 Redis 中获取购物车列表
	var carts []model.Cart
	cartJSON, err := cache.RC.Get(ctx, cartRedisKey).Result()
	if err == nil && cartJSON != "" {
		// 如果 Redis 存在缓存，解析购物车列表
		err = json.Unmarshal([]byte(cartJSON), &carts)
		if err != nil {
			zap.L().Error("购物车缓存数据解析失败", zap.String("app.cart.service", "update_cart.go"))
			return &resp.Response{
				Status: e.ERROR_UNMARSHAL_JSON,
				Msg:    e.GetMsg(e.ERROR_UNMARSHAL_JSON),
			}
		}
	} else {
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 在购物车列表中找到对应的商品并更新数量
	for i, cart := range carts {
		if cart.ProductID == service.ProductID {
			// 更新商品数量
			carts[i].Num = service.Num
			break
		}
	}

	// 更新数据库中的购物车信息
	err = db.DB.Save(&carts).Error
	if err != nil {
		zap.L().Error("数据库更新购物车信息失败", zap.String("app.cart.service", "update_cart.go"))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 更新 Redis 缓存中的购物车列表
	cartCache, err := json.Marshal(carts)
	if err != nil {
		zap.L().Error("购物车列表序列化失败", zap.String("app.cart.service", "update_cart.go"))
		return &resp.Response{
			Status: e.ERROR_UNMARSHAL_JSON,
			Msg:    e.GetMsg(e.ERROR_UNMARSHAL_JSON),
		}
	}

	// 将修改后的购物车列表重新存入 Redis
	err = cache.RC.Set(ctx, cartRedisKey, cartCache, 24*time.Hour).Err()
	if err != nil {
		zap.L().Error("更新购物车缓存失败", zap.String("app.cart.service", "update_cart.go"))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 返回最新的购物车信息
	return &resp.Response{
		Status: e.SUCCESS,
		Data:   serializer.BuildCarts(carts), // 返回更新后的购物车列表
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
