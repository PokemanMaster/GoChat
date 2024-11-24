package serializer

import (
	MFavorite "github.com/PokemanMaster/GoChat/app/favorite/model"
	MProduct "github.com/PokemanMaster/GoChat/app/product/model"
	"github.com/PokemanMaster/GoChat/common/db"
)

// Favorite 收藏序列化器
type Favorite struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"id"`

	Title    string  `json:"title"`
	Images   string  `json:"images"`
	Price    float64 `json:"price"`
	Param    string  `json:"param"`
	Saleable bool    `json:"saleable"`
	Valid    bool    `json:"valid"`
}

// BuildFavorite 序列化收藏夹
func BuildFavorite(item1 MFavorite.Favorite, item2 MProduct.ProductParam) Favorite {
	return Favorite{
		UserID:    item1.UserID,
		ProductID: item1.ProductID,
		Title:     item2.Title,
		Images:    item2.Images,
		Price:     item2.Price,
	}
}

// BuildFavorites 序列化收藏夹列表
func BuildFavorites(items []MFavorite.Favorite) (favorites []Favorite) {
	for _, item1 := range items {
		item2 := MProduct.ProductParam{}
		err := db.DB.First(&item2, item1.ProductID).Error
		if err != nil {
			continue
		}
		favorite := BuildFavorite(item1, item2)
		favorites = append(favorites, favorite)
	}
	return favorites
}
