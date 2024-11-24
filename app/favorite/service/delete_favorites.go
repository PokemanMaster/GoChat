package service

import (
	"context"
	"github.com/PokemanMaster/GoChat/app/favorite/model"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
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
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 先删除数据库中的收藏记录
	err = db.DB.Delete(&favorite).Error
	if err != nil {
		zap.L().Error("删除收藏失败", zap.String("app.favorite.service.delete_favorites", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 数据库删除成功后，删除 Redis 中的缓存
	favoriteRedisKey := "ShowFavorite_" + strconv.Itoa(int(service.UserID))
	err = cache.RC.Del(ctx, favoriteRedisKey).Err()
	if err != nil {
		zap.L().Error("删除缓存失败", zap.String("app.favorite.service.delete_favorites", ""))
	}

	// 延迟双删
	go func() {
		time.Sleep(500 * time.Millisecond)
		err = cache.RC.Del(ctx, favoriteRedisKey).Err()
		if err != nil {
			zap.L().Error("延迟删除缓存失败", zap.String("app.favorite.service.delete_favorites", ""))
		}
	}()

	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
