package model

import (
	"gorm.io/gorm"
)

// Favorite 收藏夹模型
type Favorite struct {
	gorm.Model
	UserID    uint `gorm:"not null"`                              // 用户ID
	ProductID uint `gorm:"index:idx_favorite_productId;not null"` // 商品ID
}
