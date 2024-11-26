package service

import (
	Mgroup "github.com/PokemanMaster/GoChat/server/app/group/model"
	Muser "github.com/PokemanMaster/GoChat/server/app/user/model"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
)

type GroupListsService struct {
}

// List 获取用户的群列表
func (service *GroupListsService) List(id string) resp.Response {
	// 获取关系列表
	contacts := make([]Muser.Contact, 0)
	err := db.DB.Model(&Muser.Contact{}).
		Where("owner_id = ? && type = 2", id).
		Find(&contacts).Error
	if err != nil {
		return resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	// 获取所有群id
	targetIDs := make([]uint, len(contacts))
	for i, contact := range contacts {
		targetIDs[i] = contact.TargetID
	}

	// 根据所有群id获取所有群信息
	var groups []Mgroup.GroupBasic
	err = db.DB.Model(&Mgroup.GroupBasic{}).
		Where("id IN ?", targetIDs).
		Find(&groups).Error
	if err != nil {
		return resp.Response{
			Status: e2.ERROR_DATABASE,
			Msg:    e2.GetMsg(e2.ERROR_DATABASE),
		}
	}

	return resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
		Data:   groups,
	}
}
