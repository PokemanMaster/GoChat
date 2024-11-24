package model

import (
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"gorm.io/gorm"
)

// UserAddress Address 收货地址模型
type UserAddress struct {
	gorm.Model
	UserID    uint   `gorm:"not null;comment:'客户ID';index:idx_user_id"`
	UserName  string `gorm:"type:varchar(200);not null;comment:'收货人姓名'"`
	Telephone string `gorm:"type:char(11);comment:'收货人手机号'"`
	Address   string `gorm:"type:varchar(200);not null;comment:'收货地址'"`
	Prime     bool   `gorm:"not null;comment:'是否用当前地址记录作为默认收货地址'"`
}

// GetUserAddress 获取用户地址
func (address *UserAddress) GetUserAddress(id uint) ([]UserAddress, int) {
	addresses := make([]UserAddress, 0)
	err := db.DB.Where("user_id=?", id).Order("created_at desc").Find(&addresses).Error
	if err != nil {
		logging.Info(err)
		return nil, e.ERROR_DATABASE
	}
	return addresses, e.SUCCESS
}

// SearchUserAddress 查询用户单个地址
func (address *UserAddress) SearchUserAddress(id uint) (*UserAddress, int) {
	err := db.DB.Where("id = ?", id).First(address).Error
	if err != nil {
		logging.Info("Error retrieving address:", err)
		return nil, e.ERROR_DATABASE
	}
	return address, e.SUCCESS
}
