package service

import (
	"context"
	"encoding/json"
	"github.com/PokemanMaster/GoChat/app/favorite/model"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/PokemanMaster/GoChat/resp"
	"github.com/go-redis/redis/v8"

	"strconv"
	"time"
)

// CreateFavoriteService 创建收藏服务
type CreateFavoriteService struct {
	UserID    uint
	ProductID uint
}

func (service *CreateFavoriteService) Create(ctx context.Context) *resp.Response {
	// 检查收藏是否存在
	favorite, code := model.ShowFavorite(service.UserID, service.ProductID)
	if code != e.ERROR_NOT_EXIST_FAVORITE {
		return &resp.Response{
			Status: e.ERROR_EXIST_FAVORITE,
			Msg:    e.GetMsg(e.ERROR_EXIST_FAVORITE),
		}
	}

	// 创建新的收藏
	favorite.UserID = service.UserID
	favorite.ProductID = service.ProductID
	if err := db.DB.Create(&favorite).Error; err != nil {
		logging.Info("favorite 数据库创建/更新失败", err)
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(code),
		}
	}

	// 更新缓存
	favoriteRedisKey := "ShowFavorite_" + strconv.Itoa(int(service.UserID))

	// 获取现有的收藏 JSON 数组
	existingFavoritesJSON, err := cache.RC.Get(ctx, favoriteRedisKey).Result()
	if err != nil && err != redis.Nil {
		logging.Info("从缓存获取收藏记录失败", err)
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 反序列化现有的收藏数据
	var favorites []model.Favorite
	if existingFavoritesJSON != "" {
		if err := json.Unmarshal([]byte(existingFavoritesJSON), &favorites); err != nil {
			logging.Info("反序列化收藏记录失败", err)
			return &resp.Response{
				Status: e.ERROR_DATABASE,
				Msg:    e.GetMsg(e.ERROR_DATABASE),
			}
		}
	}

	// 将新收藏追加到数组
	favorites = append(favorites, favorite)

	// 序列化更新后的收藏数组
	updatedFavoritesJSON, err := json.Marshal(favorites)
	if err != nil {
		logging.Info("序列化收藏记录失败", err)
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 将更新后的数据保存到 Redis
	err = cache.RC.Set(ctx, favoriteRedisKey, updatedFavoritesJSON, 24*time.Hour).Err()
	if err != nil {
		logging.Info("favorite 缓存更新失败", err)
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
