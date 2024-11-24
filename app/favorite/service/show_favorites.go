package service

import (
	"context"
	"encoding/json"
	"github.com/PokemanMaster/GoChat/app/favorite/model"
	"github.com/PokemanMaster/GoChat/app/favorite/serializer"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
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
		service.Limit = 12
	}
	var favorites []model.Favorite

	// 查询 redis
	favoriteJSON, err := cache.RC.Get(ctx, favoriteRedisKey).Result()
	if err == nil && favoriteJSON != "" {
		if err := json.Unmarshal([]byte(favoriteJSON), &favorites); err != nil {
			zap.L().Error("favorite 缓存数据解析失败", zap.String("app.favorite.service", "show_favorites.go"))
			return resp.Response{
				Status: e.ERROR_UNMARSHAL_JSON,
				Msg:    e.GetMsg(e.ERROR_UNMARSHAL_JSON),
			}
		}
		return resp.BuildListResponse(serializer.BuildFavorites(favorites), uint(len(favorites)))
	}

	// 如果缓存未命中，则从数据库查询
	favorites, total, code := model.ListFavorites(id, service.Limit, service.Start)
	if code != e.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 将数据库查询到的收藏数据存入 Redis
	favoritesJSON, _ := json.Marshal(favorites)
	err = cache.RC.Set(ctx, favoriteRedisKey, favoritesJSON, 24*time.Hour).Err()
	if err != nil {
		zap.L().Error("favorite 缓存创建/更新失败", zap.String("app.favorite.service", "show_favorites.go"))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(code),
		}
	}

	// 返回查询到的数据
	return resp.BuildListResponse(serializer.BuildFavorites(favorites), uint(total))
}
