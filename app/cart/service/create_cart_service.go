package service

import (
	"context"
	"encoding/json"
	MCart "github.com/PokemanMaster/GoChat/app/cart/model"
	"github.com/PokemanMaster/GoChat/app/cart/serializer"
	MProduct "github.com/PokemanMaster/GoChat/app/product/model"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/resp"
	"github.com/go-redis/redis/v8"

	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"

	"strconv"
	"time"
)

// CreateCartService 购物车创建的服务
type CreateCartService struct {
	UserID    uint
	ProductID uint
}

// Create 创建购物车
func (service *CreateCartService) Create(ctx context.Context) *resp.Response {
	// 查询商品
	productParam, code := MProduct.ShowProductParam(service.ProductID)
	if code != e.SUCCESS {
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 查询用户购物车
	cart, _, _ := MCart.ShowCart(service.UserID, service.ProductID)

	// 如果购物车不存在，创建一个新购物车
	if cart.ID == 0 {
		cart = MCart.Cart{
			UserID:    service.UserID,
			ProductID: service.ProductID,
			Num:       1,
			MaxNum:    10,
			Check:     false,
		}

		err := db.DB.Create(&cart).Error
		if err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return &resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	} else if cart.Num < cart.MaxNum {
		cart.Num++
		err := db.DB.Save(&cart).Error
		if err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return &resp.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	} else {
		return &resp.Response{
			Status: 202,
			Msg:    "超过最大上限",
		}
	}

	// 更新缓存
	cartRedisKey := "ShowCart_" + strconv.Itoa(int(service.UserID))

	// 获取现有的收藏 JSON 数组
	existingCartsJSON, err := cache.RC.Get(ctx, cartRedisKey).Result()
	if err != nil && err != redis.Nil {
		logging.Info("从缓存获取购物车记录失败", err)
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 反序列化现有的收藏数据
	var carts []MCart.Cart
	if existingCartsJSON != "" {
		if err := json.Unmarshal([]byte(existingCartsJSON), &carts); err != nil {
			logging.Info("反序列化购物车记录失败", err)
			return &resp.Response{
				Status: e.ERROR_DATABASE,
				Msg:    e.GetMsg(e.ERROR_DATABASE),
			}
		}
	}

	// 将新收藏追加到数组
	carts = append(carts, cart)

	// 序列化更新后的收藏数组
	updatedCartsJSON, err := json.Marshal(carts)
	if err != nil {
		logging.Info("序列化购物车记录失败", err)
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 将更新后的数据保存到 Redis
	err = cache.RC.Set(ctx, cartRedisKey, updatedCartsJSON, 24*time.Hour).Err()
	if err != nil {
		logging.Info("Carts 缓存更新失败", err)
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	code = e.SUCCESS

	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, productParam),
	}
}
