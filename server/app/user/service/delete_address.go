package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/user/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
)

// DeleteAddressService 购物车删除的服务
type DeleteAddressService struct {
	AddressID uint
}

// Delete 删除收货地址
func (service *DeleteAddressService) Delete() *resp.Response {
	var address model.UserAddress

	// 查询收货地址
	addressPtr, code := address.SearchUserAddress(service.AddressID)
	if code != e.SUCCESS {
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 删除收货地址并返回
	err := db.DB.Delete(addressPtr).Error
	if err != nil {
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
