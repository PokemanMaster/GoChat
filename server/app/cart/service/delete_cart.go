package service

import (
	"context"
	"github.com/PokemanMaster/GoChat/v1/server/app/cart/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
	"strconv"
)

// DeleteCartService 前端传递过来的数据
type DeleteCartService struct {
	UserID    uint
	ProductID uint
}

func (service *DeleteCartService) Delete(ctx context.Context) *resp.Response {
	// 查询购物车
	cart, code, err := model.ShowCart(service.UserID, service.ProductID)
	if code != e.SUCCESS {
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 删除购物车
	err = db.DB.Delete(&cart).Error
	if err != nil {
		zap.L().Error("删除Cart数据购物车失败", zap.String("app.cart.service", "delete_cart.go"))
		code = e.ERROR_DATABASE
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 数据库删除成功后，再删除 Redis 中的缓存
	CartRedisKey := "ShowCart_" + strconv.Itoa(int(service.UserID))
	err = cache.RC.Del(ctx, CartRedisKey).Err()
	if err != nil {
		zap.L().Error("删除Cart缓存失败", zap.String("app.cart.service", "delete_cart.go"))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	code = e.SUCCESS
	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
