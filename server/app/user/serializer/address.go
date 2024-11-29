package serializer

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/user/model"
)

// UserAddressSerialization 用户收货地址序列化器
type UserAddressSerialization struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`   // 客户ID
	UserName  string `json:"user_name"` // 收货人姓名
	Telephone string `json:"Telephone"` // 用户电话
	Address   string `json:"address"`   // 收货地址
	Prime     bool   `json:"Prime"`     // 是否用当前地址记录作为默认收货地址
}

// BuildUserAddress 收货地址购物车
func BuildUserAddress(item model.UserAddress) UserAddressSerialization {
	return UserAddressSerialization{
		ID:        item.ID,
		UserID:    item.UserID,
		UserName:  item.UserName,
		Telephone: item.Telephone,
		Address:   item.Address,
		Prime:     false,
	}
}

// BuildUserAddresses 序列化收货地址列表
func BuildUserAddresses(items []model.UserAddress) (addresses []UserAddressSerialization) {
	for _, item := range items {
		address := BuildUserAddress(item)
		addresses = append(addresses, address)
	}
	return addresses
}
