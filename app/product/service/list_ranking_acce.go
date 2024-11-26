package service

import (
	"context"
	"fmt"
	"github.com/PokemanMaster/GoChat/server/app/product/model"
	"github.com/PokemanMaster/GoChat/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/server/common/cache"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
	"go.uber.org/zap"
	"strings"
)

// ListAcceRankingService 展示配件排行的服务
type ListAcceRankingService struct {
}

// List 获取家电排行
func (service *ListAcceRankingService) List(ctx context.Context) resp.Response {
	var products []model.Product
	code := e2.SUCCESS
	// 从redis读取点击前十的视频
	pros, _ := cache.RC.ZRevRange(ctx, cache.AccessoryRank, 0, 6).Result()

	if len(pros) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(pros, ","))
		err := db.DB.Where("id in (?)", pros).Order(order).Find(&products).Error
		if err != nil {
			zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
			code := e2.ERROR_DATABASE
			return resp.Response{
				Status: code,
				Msg:    e2.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return resp.Response{
		Status: code,
		Msg:    e2.GetMsg(code),
		Data:   serializer.BuildProducts(products),
	}
}
