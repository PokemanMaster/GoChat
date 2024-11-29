package service

import (
	"context"
	MCart "github.com/PokemanMaster/GoChat/v1/server/app/cart/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/cart/serializer"
	MProduct "github.com/PokemanMaster/GoChat/v1/server/app/product/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"

	"strconv"
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
			zap.L().Error("创建购物车失败", zap.String("app.cart.service", "create_cart.go"))
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
			zap.L().Error("保存购物车失败", zap.String("app.cart.service", "create_cart.go"))
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

	// 删除缓存
	cartRedisKey := "ShowCart_" + strconv.Itoa(int(service.UserID))

	// 删除对应用户的购物车缓存
	err := cache.RC.Del(ctx, cartRedisKey).Err()
	if err != nil {
		zap.L().Error("删除缓存失败", zap.String("app.cart.service", "create_cart.go"))
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
