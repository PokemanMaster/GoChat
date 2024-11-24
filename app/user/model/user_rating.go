package model

import (
	"gorm.io/gorm"
)

// UserRating 用户评价表
type UserRating struct {
	gorm.Model
	OrderID   uint   `gorm:"type:int unsigned;not null;index;comment:'订单ID'" json:"order_id"`
	ProductID uint   `gorm:"type:int unsigned;not null;index;comment:'商品ID'" json:"product_id"`
	Img       string `gorm:"type:json;comment:'买家晒图'" json:"img"`
	Rating    uint8  `gorm:"type:tinyint unsigned;not null;comment:'评分'" json:"rating"`
	Comment   string `gorm:"type:varchar(200);comment:'评论'" json:"comment"`
}
