package service

import (
	"context"
	"github.com/PokemanMaster/GoChat/app/favorite/model"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/PokemanMaster/GoChat/resp"
	"strconv"
)

// DeleteFavoriteService 删除收藏的服务
type DeleteFavoriteService struct {
	UserID    uint
	ProductID uint
}

func (service *DeleteFavoriteService) Delete(ctx context.Context) *resp.Response {
	// 查询收藏
	favorite, code := model.ShowFavorite(service.UserID, service.ProductID)
	if code != e.SUCCESS {
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 先删除数据库中的收藏记录
	err := db.DB.Delete(&favorite).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 数据库删除成功后，再删除 Redis 中的缓存
	favoriteRedisKey := "ShowFavorite_" + strconv.Itoa(int(service.UserID))
	err = cache.RC.Del(ctx, favoriteRedisKey).Err()
	if err != nil {
		logging.Info("删除 favorite 缓存失败", err)
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
