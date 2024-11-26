package model

import (
	"github.com/PokemanMaster/GoChat/server/common/db"
	"github.com/PokemanMaster/GoChat/server/pkg/e"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UserAddress 收货地址模型
type UserAddress struct {
	gorm.Model
	UserID    uint   `gorm:"not null;index;comment:'客户ID'" json:"user_id"`
	UserName  string `gorm:"type:varchar(200);not null;comment:'收货人姓名'" json:"user_name"`
	Telephone string `gorm:"type:char(11);comment:'收货人手机号'" json:"telephone"`
	Address   string `gorm:"type:varchar(200);not null;comment:'收货地址'" json:"address"`
	Prime     bool   `gorm:"not null;comment:'是否为默认收货地址'" json:"prime"`
}

// GetUserAddress 获取用户地址
func (address *UserAddress) GetUserAddress(id uint) ([]UserAddress, int) {
	addresses := make([]UserAddress, 0)
	err := db.DB.Where("user_id=?", id).Order("created_at desc").Find(&addresses).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		return nil, e.ERROR_DATABASE
	}
	return addresses, e.SUCCESS
}

// SearchUserAddress 查询用户单个地址
func (address *UserAddress) SearchUserAddress(id uint) (*UserAddress, int) {
	err := db.DB.Where("id = ?", id).First(address).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		return nil, e.ERROR_DATABASE
	}
	return address, e.SUCCESS
}
