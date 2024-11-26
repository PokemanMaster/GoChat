package model

import (
	"github.com/PokemanMaster/GoChat/server/common/db"
	"github.com/PokemanMaster/GoChat/server/pkg/e"
)

// OrderDetail 订单详情表
type OrderDetail struct {
	OrderID     uint    `gorm:"type:int unsigned;not null;comment:'订单ID';primary_key"`
	ProductID   uint    `gorm:"type:int unsigned;not null;comment:'商品ID';primary_key"`
	Price       float64 `gorm:"type:decimal(10,2) unsigned;not null;comment:'原价格'"`
	ActualPrice float64 `gorm:"type:decimal(10,2) unsigned;not null;comment:'实际购买价格'"`
	Num         uint    `gorm:"type:int unsigned;not null;comment:'购买数量'"`
}

// ToMap 实现 ToMap 方法
func (orderDetail OrderDetail) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"order_id":     orderDetail.OrderID,
		"warehouse_id": 1,
		"product_id":   orderDetail.ProductID,
		"reduce_num":   orderDetail.Num,
	}
}

func ShowOrderDetail(id string) (OrderDetail, int, error) {
	var orderDetail OrderDetail
	err := db.DB.Where("order_id=?", id).First(&orderDetail).Error
	if err != nil {
		return orderDetail, e.ERROR_DATABASE, err
	}
	return orderDetail, e.SUCCESS, nil
}
