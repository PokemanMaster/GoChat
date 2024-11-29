package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/user/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/user/serializer"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/mid"

	"github.com/PokemanMaster/GoChat/v1/server/resp"
)

// CreateAddressService 收货地址创建的服务
type CreateAddressService struct {
	UserID    uint
	UserName  string
	Telephone string
	Address   string
}

// Create 用户创建收货地址，同时展示自己已经创建过的地址
func (service *CreateAddressService) Create() *resp.Response {
	address := model.UserAddress{
		UserID:    service.UserID,
		UserName:  service.UserName,
		Telephone: service.Telephone,
		Address:   service.Address,
	}
	code := e.SUCCESS

	// 检查电话
	code = mid.TelephoneNumberIsTure(service.Telephone)
	if code != e.SUCCESS {
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建收货地址
	err := db.DB.Create(&address).Error
	if err != nil {
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 查询收货地址返回
	addresses, code := address.GetUserAddress(service.UserID)
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
