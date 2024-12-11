package model

import (
	"context"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

// ProductParam 商品参数 (Sku)
type ProductParam struct {
	gorm.Model
	ProductID     uint    `gorm:"type:int;not null;index:idx_product_id;comment:'商品ID'"`
	Image         string  `gorm:"type:varchar(200);not null;comment:'商品的次图'"`
	Price         float64 `gorm:"type:decimal(10,2) unsigned;not null;comment:'商品价格'"`
	DiscountPrice float64 `gorm:"type:decimal(10,2) unsigned;comment:'商品折扣价格'"`
	Stock         uint    `gorm:"type:int;not null;comment:'库存数量'"`
	SoldCount     uint    `gorm:"type:int;not null;default:0;comment:'已售数量'"`
	Weight        float64 `gorm:"type:decimal(10,2);comment:'商品重量'"`
	Color         string  `gorm:"type:varchar(20);comment:'颜色'"`
	Size          string  `gorm:"type:varchar(20);comment:'尺寸'"`
	Saleable      bool    `gorm:"not null;default:true;comment:'是否上架'"`
}

// SearchProductParam 搜索商品
func SearchProductParam(search string) ([]ProductParam, error) {
	var productParam []ProductParam
	err := db.DB.Where("title LIKE ?", "%"+search+"%").Find(&productParam).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		return productParam, err
	}
	return productParam, nil
}

// ShowProductParam 获取商品参数
func ShowProductParam(productId uint) (ProductParam, int) {
	var productParam ProductParam
	err := db.DB.Where("product_id = ?", productId).First(&productParam).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		return productParam, e.ERROR_DATABASE
	}
	return productParam, e.SUCCESS
}

// View 获取点击数
func (productParam *ProductParam) View(ctx context.Context) uint64 {
	countStr, _ := cache.RC.Get(ctx, cache.ProductViewKey(productParam.ProductID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 视频游览
func (productParam *ProductParam) AddView(ctx context.Context) {
	// 增加视频点击数
	cache.RC.Incr(ctx, cache.ProductViewKey(productParam.ProductID))
	// 增加排行点击数
	cache.RC.ZIncrBy(ctx, cache.RankKey, 1, strconv.Itoa(int(productParam.ProductID)))
}

// AddElecRank 增加家电排行点击数
func (productParam *ProductParam) AddElecRank(ctx context.Context) {
	// 增加家电排行点击数
	cache.RC.ZIncrBy(ctx, cache.ElectricalRank, 1, strconv.Itoa(int(productParam.ProductID)))
}

// AddAcceRank 增加配件排行点击数
func (productParam *ProductParam) AddAcceRank(ctx context.Context) {
	// 增加配件排行点击数
	cache.RC.ZIncrBy(ctx, cache.AccessoryRank, 1, strconv.Itoa(int(productParam.ProductID)))
}
