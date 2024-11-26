package serializer

import (
	MCart "github.com/PokemanMaster/GoChat/server/app/cart/model"
	MProduct "github.com/PokemanMaster/GoChat/server/app/product/model"
	"github.com/PokemanMaster/GoChat/server/common/db"
)

// CartSerialization 购物车序列化器
type CartSerialization struct {
	ID        uint `json:"id"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Num       uint `json:"num"`
	MaxNum    uint `json:"max_num"`
	Check     bool `json:"check"`

	Title  string  `json:"title"`
	Images string  `json:"images"`
	Price  float64 `json:"price"`
}

// BuildCart 序列化购物车
func BuildCart(item1 MCart.Cart, item2 MProduct.ProductParam) CartSerialization {
	return CartSerialization{
		ID:        item1.ID,
		UserID:    item1.UserID,
		ProductID: item1.ProductID,
		Num:       item1.Num,
		MaxNum:    item1.MaxNum,
		Check:     item1.Check,
		Title:     item2.Title,
		Images:    item2.Images,
		Price:     item2.Price,
	}
}

// BuildCarts 序列化购物车列表
func BuildCarts(items []MCart.Cart) (carts []CartSerialization) {
	for _, item1 := range items {
		item2 := MProduct.ProductParam{}
		err := db.DB.First(&item2, item1.ProductID).Error
		if err != nil {
			continue
		}
		cart := BuildCart(item1, item2)
		carts = append(carts, cart)
	}
	return carts
}
