package service

import (
	"context"
	"errors"
	"github.com/PokemanMaster/GoChat/server/app/favorite/model"
	"github.com/PokemanMaster/GoChat/server/common/cache"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
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
				Status: e2.ERROR_DATABASE,
				Msg:    e2.GetMsg(e2.ERROR_DATABASE),
			}
		}
	} else {
		zap.L().Info("favorite已存在", zap.String("app.favorite.service.create_favorites", ""))
		return &resp.Response{
			Status: e2.ERROR_EXIST_FAVORITE,
			Msg:    e2.GetMsg(e2.ERROR_EXIST_FAVORITE),
		}
	}

	// 更新数据库
	favorite.UserID = service.UserID
	favorite.ProductID = service.ProductID
	if err = db.DB.Create(&favorite).Error; err != nil {
		zap.L().Error("数据创建失败", zap.String("app.favorite.service.create_favorites", ""))
		return &resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	// 删除缓存
	favoriteRedisKey := "ShowFavorite_" + strconv.Itoa(int(service.UserID))
	go func() {
		time.Sleep(500 * time.Millisecond)
		err = cache.RC.Del(ctx, favoriteRedisKey).Err()
		if err != nil {
			zap.L().Error("延迟删除缓存失败", zap.String("app.favorite.service.create_favorites", ""))
		}
	}()

	return &resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
	}
}
