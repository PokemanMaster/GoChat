package service

import (
	"github.com/PokemanMaster/GoChat/app/user/model"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
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
			Data:   []model.UserBasic{},
		}
	}

	var users []model.UserBasic
	query := db.DB.Model(&model.UserBasic{}).
		Where("id IN ?", contactIDs).
		Where("name LIKE ?", "%"+service.FriendName+"%")

	err = query.Find(&users).Error
	if err != nil {
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
