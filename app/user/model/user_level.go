package model

import "gorm.io/gorm"

// UserLevel Level 会员等级表
type UserLevel struct {
	gorm.Model
	Level    string  `gorm:"type:varchar(200);not null;comment:'等级'"`  // 等级名称；如金牌、银牌
	Discount float64 `gorm:"type:decimal(10,2);not null;comment:'折扣'"` // 折扣，享受的折扣
}
