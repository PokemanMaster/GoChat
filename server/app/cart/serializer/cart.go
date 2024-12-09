package serializer

import (
	MCart "github.com/PokemanMaster/GoChat/v1/server/app/cart/model"
	MProduct "github.com/PokemanMaster/GoChat/v1/server/app/product/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"log"
)

// CartSerialization 购物车序列化器
type CartSerialization struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Num       uint    `json:"num"`
	Check     bool    `json:"check"`
	Name      string  `json:"name"`
	Image     string  `json:"image"`
	Price     float64 `json:"price"`
}

// BuildCart 序列化购物车
func BuildCart(item1 MCart.Cart, item2 MProduct.ProductParam, item3 MProduct.Product) CartSerialization {
	return CartSerialization{
		ID:        item1.ID,
		UserID:    item1.UserID,
		ProductID: item1.ProductID,
		Name:      item3.Name,
		Num:       item1.Num,
		Check:     item1.Check,
		Image:     item2.Image,
		Price:     item2.Price,
	}
}

// BuildCarts 优化查询逻辑，减少数据库查询次数
func BuildCarts(items []MCart.Cart) (carts []CartSerialization) {
	for _, item1 := range items {
		item2 := MProduct.ProductParam{}
		item3 := MProduct.Product{}

		// 获取 ProductParam 和 Product 信息
		err := db.DB.First(&item2, item1.ProductID).Error
		if err != nil {
			log.Printf("Error fetching ProductParam for ProductID %d: %v\n", item1.ProductID, err)
			continue
		}

		err = db.DB.First(&item3, item1.ProductID).Error
		if err != nil {
			log.Printf("Error fetching Product for ProductID %d: %v\n", item1.ProductID, err)
			continue
		}

		// 调用 BuildCart 序列化购物车项
		cart := BuildCart(item1, item2, item3)
		carts = append(carts, cart)
	}

	return carts
}
