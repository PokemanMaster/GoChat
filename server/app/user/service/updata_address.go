package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/user/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/user/serializer"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
)

// UpdateAddressService 收货地址修改的服务
type UpdateAddressService struct {
	ID        uint
	UserID    uint
	UserName  string
	Telephone string
	Address   string
	Prime     bool
}

// Update 修改购物车信息
func (service *UpdateAddressService) Update() *resp.Response {
	address := model.UserAddress{
		UserID:    service.UserID,
		UserName:  service.UserName,
		Telephone: service.Telephone,
		Address:   service.Address,
		Prime:     service.Prime,
	}

	address.ID = service.ID
	code := e.SUCCESS

	// 修改收货地址
	err := db.DB.Save(&address).Error
	if err != nil {
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 查询收货地址并返回
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
