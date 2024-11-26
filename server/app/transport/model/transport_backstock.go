package model

import "time"

// TransportBackstock Backstock 退货表
type TransportBackstock struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;not null;comment:'主键'"`
	OrderID     uint      `gorm:"type:int unsigned;not null;index:idx_order_id;comment:'订单ID'"`
	ProductID   uint      `gorm:"type:int unsigned;not null;comment:'商品ID'"`
	Reason      string    `gorm:"type:varchar(200);not null;comment:'退货原因'"`
	QAID        uint      `gorm:"type:int unsigned;not null;index:idx_qa_id;comment:'质检员ID'"`
	Payment     float64   `gorm:"type:decimal(10,2) unsigned;not null;comment:'退款金额'"`
	PaymentType uint8     `gorm:"type:tinyint unsigned;not null;comment:'退款方式：1借记卡、2信用卡、3微信、4支付宝、5现金'"`
	Status      uint8     `gorm:"type:tinyint unsigned;not null;index:idx_status;comment:'状态：1退货成功、2无法退货'"`
	CreateTime  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'添加时间'"`
}
