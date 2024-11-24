package model

import (
	"context"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"strconv"
	"time"
)

// ProductParam 商品参数 (Sku)
type ProductParam struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null;comment:'主键'"`
	ProductID uint   `gorm:"not null;comment:'产品ID';index:idx_product_id"`
	Title     string `gorm:"type:varchar(200);not null;comment:'标题'"`
	//  desc: 商品描述图
	//	facade：商品展示图
	//{"desc": \["http://127.0.0.1/1.jpg", "http://127.0.0.1/2.jpg"\], "facade": \["http://127.0.0.1/3.jpg", "http://127.0.0.1/4.jpg"\]}
	Images string `gorm:"type:json;comment:'商品图片'"`
	//需要说明的是，当促销时，会有促销价格，需要再多一个额外的价格字段吗？
	//这个要看业务场景，比如我们做的新零售系统，以会员制，每个会员等级享受折扣不同，就不适合都定义在商品表中，
	//另外在客户部中定义「会员等级」字段。客户浏览商品的时候看到的价格就是该客户端的会员等级价格。
	Price          float64   `gorm:"type:decimal(10,2) unsigned;not null;comment:'价格'"`
	Param          string    `gorm:"type:json;not null;comment:'参数'"`            //{"CPU": "骁龙855", "内存": "128", "电池": 4000, "运存": 8, "屏幕尺寸": 6.39}
	Saleable       bool      `gorm:"not null;index:idx_saleable;comment:'是否上架'"` // 1 表示上架，0 表示未上架
	Valid          bool      `gorm:"not null;index:idx_valid;comment:'是否有效'"`
	CreateTime     time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_create_time;comment:'添加时间'"`
	LastUpdateTime time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index:idx_last_update_time;comment:'最后修改时间'"`
}

// SearchProductParam 搜索商品
func SearchProductParam(search string) ([]ProductParam, int, error) {
	var productParam []ProductParam
	err := db.DB.Where("title LIKE ?", "%"+search+"%").Find(&productParam).Error
	if err != nil {
		logging.Info(err)
		return productParam, e.ERROR_DATABASE, err
	}
	return productParam, e.SUCCESS, err
}

// ShowProductParam 获取商品参数
func ShowProductParam(productId uint) (ProductParam, int) {
	var productParam ProductParam
	err := db.DB.Where("product_id = ?", productId).First(&productParam).Error
	if err != nil {
		logging.Info(err)
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
