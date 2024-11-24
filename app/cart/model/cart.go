package model

import (
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Cart 购物车
type Cart struct {
	gorm.Model
	UserID    uint // 用户id
	ProductID uint // 商品id
	Num       uint // 商品数量
	MaxNum    uint // 订单号
	Check     bool // 是否被选中
}

// ShowCart 获取用户的购物车
func ShowCart(UserId, ProductId uint) (Cart, int, error) {
	var cart Cart
	err := db.DB.Where("user_id = ? AND product_id = ?", UserId, ProductId).First(&cart).Error
	if err != nil {
		zap.L().Info("获取用户购物车失败", zap.String("app.cart.model", "cart.go"))
		return cart, e.ERROR_DATABASE, err
	}
	return cart, e.SUCCESS, err
}

// ListCart 获取用户的购物车列表
func ListCart(id string) ([]Cart, int) {
	var cart []Cart
	err := db.DB.Where("user_id = ?", id).Find(&cart).Error
	if err != nil {
		zap.L().Info("获取用户的购物车列表失败", zap.String("app.cart.model", "cart.go"))
		return cart, e.ERROR_DATABASE
	}
	return cart, e.SUCCESS
}
