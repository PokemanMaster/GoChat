package model

import (
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"gorm.io/gorm"
)

// Favorite 收藏夹模型
type Favorite struct {
	gorm.Model
	UserID    uint `gorm:"not null"`                              // 用户ID
	ProductID uint `gorm:"index:idx_favorite_productId;not null"` // 商品ID
}

// ShowFavorite 获取用户收藏
func ShowFavorite(userId, productId uint) (Favorite, int) {
	var favorite Favorite
	err := db.DB.Where("user_id = ? AND product_id = ?", userId, productId).First(&favorite).Error
	if err != nil {
		return favorite, e.ERROR_DATABASE
	}
	return favorite, e.SUCCESS
}

// ListFavorites 获取用户收藏列表，分页查询
func ListFavorites(id string, Limit int, Start int) ([]Favorite, int64, int) {
	var favorites []Favorite
	var total int64
	if err := db.DB.Model(&favorites).Where("user_id=?", id).Count(&total).Error; err != nil {
		logging.Info(err)
		return favorites, total, e.ERROR_DATABASE
	}
	err := db.DB.Where("user_id=?", id).Limit(Limit).Offset(Start).Find(&favorites).Error
	if err != nil {
		logging.Info(err)
		return favorites, total, e.ERROR_DATABASE
	}
	return favorites, total, e.SUCCESS
}
