package model

import (
	"context"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"
	"gorm.io/gorm"

	"strconv"
)

// Product 产品表结构 (Spu)
type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(200);not null;index:idx_name;comment:'商品名称'"`
	CategoryID  uint   `gorm:"type:int unsigned;not null;comment:'商品类别ID'"`
	BrandID     uint   `gorm:"type:int unsigned;comment:'品牌ID'"`
	Image       string `gorm:"type:varchar(200);not null;comment:'商品的主图'"`
	Description string `gorm:"type:json;comment:'商品描述图片（JSON数组形式存储）'"`
	Rating      uint   `gorm:"type:int unsigned;not null;comment:'商品评分'"`
	Saleable    bool   `gorm:"not null;comment:'是否上架'"`
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
