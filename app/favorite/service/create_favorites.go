package service

import (
	"context"
	"errors"
	"github.com/PokemanMaster/GoChat/app/favorite/model"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
	"go.uber.org/zap"
	"gorm.io/gorm"

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
	var favorite model.Favorite
	err := db.DB.Where("user_id = ? AND product_id = ?", service.UserID, service.ProductID).First(&favorite).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Info("favorite不存在", zap.String("app.favorite.service.create_favorites", ""))
		} else {
			zap.L().Info("数据查询失败", zap.String("app.favorite.service.create_favorites", err.Error()))
			return &resp.Response{
				Status: e.ERROR_DATABASE,
				Msg:    e.GetMsg(e.ERROR_DATABASE),
			}
		}
	} else {
		zap.L().Info("favorite已存在", zap.String("app.favorite.service.create_favorites", ""))
		return &resp.Response{
			Status: e.ERROR_EXIST_FAVORITE,
			Msg:    e.GetMsg(e.ERROR_EXIST_FAVORITE),
		}
	}

	// 创建新的收藏
	favorite.UserID = service.UserID
	favorite.ProductID = service.ProductID
	if err := db.DB.Create(&favorite).Error; err != nil {
		zap.L().Error("favorite 数据库创建/更新失败", zap.String("app.favorite.service.create_favorites", ""))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// redis key
	favoriteRedisKey := "ShowFavorite_" + strconv.Itoa(int(service.UserID))

	// 删除缓存
	err = cache.RC.Del(ctx, favoriteRedisKey).Err()
	if err != nil {
		zap.L().Error("删除缓存失败", zap.String("app.favorite.service.create_favorites", ""))
	}

	// 更新数据库
	favorite.UserID = service.UserID
	favorite.ProductID = service.ProductID
	if err = db.DB.Create(&favorite).Error; err != nil {
		zap.L().Error("favorite 数据库创建失败", zap.String("app.favorite.service.create_favorites", ""))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 删除缓存
	go func() {
		time.Sleep(500 * time.Millisecond)
		err = cache.RC.Del(ctx, favoriteRedisKey).Err()
		if err != nil {
			zap.L().Error("延迟删除缓存失败", zap.String("app.favorite.service.create_favorites", ""))
		}
	}()

	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
