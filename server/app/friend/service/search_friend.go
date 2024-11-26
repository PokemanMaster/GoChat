package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/user/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
)

type SearchFriendService struct {
	UserID     uint
	FriendName string
}

// Search 查找好友
func (service *SearchFriendService) Search() *resp.Response {
	var contactIDs []uint
	err := db.DB.Model(&model.Contact{}).
		Where("owner_id = ? AND type = ?", service.UserID, 1).
		Pluck("target_id", &contactIDs).Error
	if err != nil {
		zap.L().Info("查询好友列表失败", zap.String("app.friend.service.search_friend", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 如果没有好友，直接返回空结果
	if len(contactIDs) == 0 {
		return &resp.Response{
			Status: e.SUCCESS,
			Msg:    e.GetMsg(e.SUCCESS),
			Data:   []model.User{},
		}
	}

	var users []model.User
	query := db.DB.Model(&model.User{}).
		Where("id IN ?", contactIDs).
		Where("user_name LIKE ?", "%"+service.FriendName+"%")

	err = query.Find(&users).Error
	if err != nil {
		zap.L().Info("搜索好友失败", zap.String("app.friend.service.search_friend", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data:   users,
	}
}
