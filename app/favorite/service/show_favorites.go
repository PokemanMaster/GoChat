package service

import (
	"context"
	"encoding/json"
	"github.com/PokemanMaster/GoChat/app/favorite/model"
	"github.com/PokemanMaster/GoChat/app/favorite/serializer"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"time"
)

// ShowFavoritesService 展示收藏夹详情的服务
type ShowFavoritesService struct {
	Limit int
	Start int
}

func (service *ShowFavoritesService) Show(ctx context.Context, id string) resp.Response {
	favoriteRedisKey := "ShowFavorite_" + id
	if service.Limit <= 0 {
		service.Limit = 17
	}
	var favorites []model.Favorite

	// 查询缓存数据
	favoriteJSON, err := cache.RC.Get(ctx, favoriteRedisKey).Result()
	if err == nil && favoriteJSON != "" {
		err = json.Unmarshal([]byte(favoriteJSON), &favorites)
		if err != nil {
			zap.L().Error("反序列化失败", zap.String("app.favorite.service.show_favorites", ""))
			return resp.Response{
				Status: e.ERROR_UNMARSHAL_JSON,
				Msg:    e.GetMsg(e.ERROR_UNMARSHAL_JSON),
			}
		}
		return resp.BuildResponseTotal(serializer.BuildFavorites(favorites), uint(len(favorites)))
	} else if err != redis.Nil {
		zap.L().Error("查询异常记录日志", zap.String("app.favorite.service.show_favorites", err.Error()))
	}

	// 查询数据库总数
	var total int64
	err = db.DB.Model(&model.Favorite{}).Where("user_id = ?", id).Count(&total).Error
	if err != nil {
		zap.L().Error("数据库查询总数失败", zap.String("app.favorite.service.show_favorites", err.Error()))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 查询数据库分页数据
	err = db.DB.Where("user_id = ?", id).Limit(service.Limit).Offset(service.Start).Find(&favorites).Error
	if err != nil {
		zap.L().Error("数据库查询分页数据失败", zap.String("app.favorite.service.show_favorites", err.Error()))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 将分页数据存入 Redis
	favoritesJSON, _ := json.Marshal(favorites)
	err = cache.RC.Set(ctx, favoriteRedisKey, favoritesJSON, 24*time.Hour).Err()
	if err != nil {
		zap.L().Error("缓存创建失败", zap.String("app.favorite.service.show_favorites", ""))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 返回查询到的数据
	return resp.BuildResponseTotal(serializer.BuildFavorites(favorites), uint(total))
}
