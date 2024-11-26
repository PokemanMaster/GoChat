package service

import (
	model2 "github.com/PokemanMaster/GoChat/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
	"go.uber.org/zap"
)

type SearchFriendService struct {
	UserID     uint
	FriendName string
}

// Search 查找好友
func (service *SearchFriendService) Search() *resp.Response {
	var contactIDs []uint
	err := db.DB.Model(&model2.Contact{}).
		Where("owner_id = ? AND type = ?", service.UserID, 1).
		Pluck("target_id", &contactIDs).Error
	if err != nil {
		zap.L().Info("查询好友列表失败", zap.String("app.friend.service.search_friend", err.Error()))
		return &resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	// 如果没有好友，直接返回空结果
	if len(contactIDs) == 0 {
		return &resp.Response{
			Status: e2.SUCCESS,
			Msg:    e2.GetMsg(e2.SUCCESS),
			Data:   []model2.User{},
		}
	}

	var users []model2.User
	query := db.DB.Model(&model2.User{}).
		Where("id IN ?", contactIDs).
		Where("user_name LIKE ?", "%"+service.FriendName+"%")

	err = query.Find(&users).Error
	if err != nil {
		zap.L().Info("搜索好友失败", zap.String("app.friend.service.search_friend", err.Error()))
		return &resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	return &resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
		Data:   users,
	}
}
