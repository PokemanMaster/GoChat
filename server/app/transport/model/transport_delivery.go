package model

import (
	"github.com/PokemanMaster/GoChat/server/server/common/db"
	"github.com/PokemanMaster/GoChat/server/server/pkg/e"
	"time"
)

// TransportDelivery Delivery 快递表
type TransportDelivery struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;not null;comment:'主键'"`
	OrderID     uint      `gorm:"type:int unsigned;not null;index:idx_order_id;comment:'订单ID'"`
	ProductID   uint      `gorm:"type:int unsigned;not null;comment:'商品ID'"`
	QAID        uint      `gorm:"type:int unsigned;not null;index:idx_qa_id;comment:'质检员ID'"`
	DEID        uint      `gorm:"type:int unsigned;not null;index:idx_de_id;comment:'发货员ID'"`
	PostID      uint      `gorm:"type:int unsigned;not null;index:idx_postid;comment:'快递单号'"`
	Price       float64   `gorm:"type:decimal(10,2) unsigned;not null;comment:'快递费'"`
	AddressID   uint      `gorm:"type:int unsigned;not null;index:idx_address_id;comment:'收货地址ID'"`
	WarehouseID uint      `gorm:"type:int unsigned;not null;index:idx_warehouse_id;comment:'发货仓库ID'"`
	ECP         uint8     `gorm:"type:tinyint unsigned;not null;index:idx_ecp;comment:'快递公司编号'"`
	CreateTime  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'添加时间'"`
}

func ShowTransportDelivery(id string) (TransportDelivery, int, error) {
	var delivery TransportDelivery
	err := db.DB.Where("order_id=?", id).First(&delivery).Error
	if err != nil {
		return delivery, e.ERROR_DATABASE, err
	}
	return delivery, e.SUCCESS, nil
}
