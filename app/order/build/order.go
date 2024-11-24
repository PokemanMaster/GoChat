package build

import (
	MOrder "github.com/PokemanMaster/GoChat/app/order/model"
	MProduct "github.com/PokemanMaster/GoChat/app/product/model"
	MTransport "github.com/PokemanMaster/GoChat/app/transport/model"
	"github.com/PokemanMaster/GoChat/common/db"
)

// UserOrdersSerialization 某个用户的所有订单
type UserOrdersSerialization struct {
	ID          uint    `json:"id"`
	Code        string  `json:"code"`
	UserID      uint    `json:"user_id"`
	Status      uint8   `json:"status"`
	CreateTime  int64   `json:"created_time"`
	ProductID   uint    `json:"product_id"`
	Price       float64 `json:"price"`
	ActualPrice float64 `json:"actualPrice"`
	Num         uint    `json:"num"`
	Title       string  `json:"title"`
	Images      string  `json:"images"`
}

// ResUserOrder 序列化某个用户的所有订单
func ResUserOrder(item1 MOrder.Order, item2 MOrder.OrderDetail, item3 MProduct.ProductParam) UserOrdersSerialization {
	return UserOrdersSerialization{
		ID:          item1.ID,
		Code:        item1.Code,
		UserID:      item1.UserID,
		Status:      item1.Status, // 状态：1未付款、2已付款、3已发货、4已签收
		ProductID:   item2.ProductID,
		Num:         item2.Num,
		Price:       item2.Price,
		ActualPrice: item2.ActualPrice,
		Images:      item3.Images,
		Title:       item3.Title,
	}
}

// ResUserOrders 序列化订单
func ResUserOrders(items []MOrder.Order) (orders []UserOrdersSerialization) {
	for _, item1 := range items {
		item2 := MOrder.OrderDetail{}
		item3 := MProduct.ProductParam{}
		err := db.DB.First(&item2, item1.ID).Error
		err = db.DB.First(&item3, item2.ProductID).Error
		if err != nil {
			continue
		}
		order := ResUserOrder(item1, item2, item3)
		orders = append(orders, order)
	}
	return orders
}

// OrderSerialization 某个用户的某个订单
type OrderSerialization struct {
	ID          uint    `json:"id"`
	Code        string  `json:"code"`
	Type        uint8   `json:"type"`
	ShopID      uint    `json:"shop_id"`
	UserID      uint    `json:"user_id"`
	Amount      float64 `json:"amount"`
	PaymentType uint8   `json:"paymentType"`
	Status      uint8   `json:"status"`
	Postage     float64 `json:"postage"`
	Weight      uint    `json:"weight"`
	CreateTime  int64   `json:"created_time"`

	OrderID     uint    `json:"orderID"`
	ProductID   uint    `json:"productID"`
	Price       float64 `json:"price"`
	ActualPrice float64 `json:"actualPrice"`
	Num         uint    `json:"num"`

	Title    string `json:"title"`
	Images   string `json:"images"`
	Param    string `json:"param"`
	Saleable bool   `json:"saleable"`
	Valid    bool   `json:"valid"`

	AddressID uint `json:"address_id"`
}

// ResOrder 序列化订单
func ResOrder(item1 MOrder.Order, item2 MOrder.OrderDetail, item3 MProduct.ProductParam, item4 MTransport.TransportDelivery) OrderSerialization {
	return OrderSerialization{
		ID:          item1.ID,
		Code:        item1.Code,
		UserID:      item1.UserID,
		ProductID:   item2.ProductID,
		Num:         item2.Num,
		Price:       item2.Price,
		ActualPrice: item2.ActualPrice,
		Title:       item3.Title,
		Param:       item3.Param,
		Images:      item3.Images,
		AddressID:   item4.AddressID,
	}
}
