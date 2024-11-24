package model

import (
	"context"
	"github.com/PokemanMaster/GoChat/common/cache"
	"gorm.io/gorm"

	"strconv"
)

// Product 产品表结构 (Spu)
type Product struct {
	// 比如，去年有一件风衣卖得很火，今年不太行了。
	// 商家不打算卖了，就删除该产品。但是还关联了那么多的订单，我们不能物理删除，就使用它逻辑删除。
	gorm.Model
	Title      string `gorm:"type:varchar(200);not null;comment:'标题'"`
	SubTitle   string `gorm:"type:varchar(200);comment:'副标题'"`
	CategoryID uint   `gorm:"type:int unsigned;not null;index:idx_category_id;comment:'分类ID'"`
	BrandID    uint   `gorm:"type:int unsigned;index:idx_brand_id;comment:'品牌ID'"` // 一些散装的可能没有品牌，比如花生之类的
	SpgID      uint   `gorm:"type:int unsigned;index:idx_spg_id;comment:'品类ID'"`   // 电子产品可能包含：电脑、电视等等，相当于二级索引
	Saleable   bool   `gorm:"not null;index:idx_saleable;comment:'是否上架'"`
	Valid      bool   `gorm:"not null;index:idx_valid;comment:'是否有效'"`
}

// View 获取点击数
func (product *Product) View(ctx context.Context) uint64 {
	countStr, _ := cache.RC.Get(ctx, cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 视频游览
func (product *Product) AddView(ctx context.Context) {
	// 增加视频点击数
	cache.RC.Incr(ctx, cache.ProductViewKey(product.ID))
	// 增加排行点击数
	cache.RC.ZIncrBy(ctx, cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}

// AddElecRank 增加家电排行点击数
func (product *Product) AddElecRank(ctx context.Context) {
	// 增加家电排行点击数
	cache.RC.ZIncrBy(ctx, cache.ElectricalRank, 1, strconv.Itoa(int(product.ID)))
}

// AddAcceRank 增加配件排行点击数
func (product *Product) AddAcceRank(ctx context.Context) {
	// 增加配件排行点击数
	cache.RC.ZIncrBy(ctx, cache.AccessoryRank, 1, strconv.Itoa(int(product.ID)))
}
