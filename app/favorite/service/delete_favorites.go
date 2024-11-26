package service

import (
	"context"
	"github.com/PokemanMaster/GoChat/server/app/favorite/model"
	"github.com/PokemanMaster/GoChat/server/common/cache"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// DeleteFavoriteService 删除收藏的服务
type DeleteFavoriteService struct {
	UserID    uint
	ProductID uint
}

func (service *DeleteFavoriteService) Delete(ctx context.Context) *resp.Response {
	var favorite model.Favorite
	// 查询数据库的收藏
	err := db.DB.Where("user_id = ? AND product_id = ?", service.UserID, service.ProductID).First(&favorite).Error
	if err != nil {
		zap.L().Error("查询收藏失败", zap.String("app.favorite.service.delete_favorites", err.Error()))
		return &resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	// 数据库删除成功后，删除 Redis 中的缓存
	favoriteRedisKey := "ShowFavorite_" + strconv.Itoa(int(service.UserID))
	err = cache.RC.Del(ctx, favoriteRedisKey).Err()
	if err != nil {
		zap.L().Error("删除缓存失败", zap.String("app.favorite.service.delete_favorites", ""))
	}

	go func() {
		time.Sleep(500 * time.Millisecond)
		err = cache.RC.Del(ctx, favoriteRedisKey).Err()
		if err != nil {
			zap.L().Error("删除缓存失败", zap.String("app.favorite.service.delete_favorites", ""))
		}
	}()

	return &resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
	}
}
