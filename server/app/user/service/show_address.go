package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/user/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/user/serializer"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"strconv"
)

// ShowAddressesService 展示收货地址的服务
type ShowAddressesService struct{}

// Show 展示用户的收货地址
func (service *ShowAddressesService) Show(id string) *resp.Response {
	var address model.UserAddress

	// 查询收货地址返回
	userId, _ := strconv.Atoi(id)
	addresses, code := address.GetUserAddress(uint(userId))
	if code != e.SUCCESS {
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUserAddresses(addresses),
	}
}
