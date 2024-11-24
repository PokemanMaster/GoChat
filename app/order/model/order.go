package model

import (
	"context"
	"fmt"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/go-redis/redis/v8"

	"gorm.io/gorm"
	"math/rand"
	"os"

	"strconv"
	"time"
)

// 一张订单中可以包含多个商品记录，可不可以用 JSON 存储这些商品信息？不适合
//因为 ：MySQL 5.7+ 引入的 JSON 字段适合存储数据，不适合检索数据。
//这里存储的数据只是用来页面展示，而 不用来做搜索条件，在前面设计的表中，很少对字符串的字段做索引。

// Order 订单
type Order struct {
	gorm.Model
	// 流水号尾部以字母 A 结尾表示液体，B 结尾表示易碎品，等。还可以包含订单的日期、和时间，消费小票打印出来后，
	// 阅读流水号可以知道是哪一类的商品、什么时候生成的订单、是否加急发货等。这些规则由业务人员去制定。
	Code        string  `gorm:"type:varchar(200);not null;comment:'流水号';uniqueIndex:idx_code"`
	Type        uint8   `gorm:"type:tinyint unsigned;not null;comment:'订单类型：1实体销售，2网络销售';index:idx_type"` // 如为 1：表示线下销售商品，为 2 为线上销售的商品
	ShopID      uint    `gorm:"comment:'零售店ID';index:idx_shop_id"`                                        // 如果线上卖出的，可以为空
	UserID      uint    `gorm:"comment:'会员ID';index:idx_user_id"`                                         // 主要用来关联会员等级等。如果用户在线下购买的商品，不是会员可以为空
	Amount      float64 `gorm:"type:decimal(10,2) unsigned;not null;comment:'总金额'"`
	PaymentType uint8   `gorm:"type:tinyint unsigned;not null;comment:'支付方式：1借记卡、2信用卡、3微信、4支付宝、5现金'"`             // 比如是借记卡付款、信用卡、微信支付、现今支付等等
	Status      uint8   `gorm:"type:tinyint unsigned;not null;comment:'状态：1未付款、2已付款、3已发货、4已签收';index:idx_status"` // 如：未付款、已付款、已发货、已签收等
	Postage     float64 `gorm:"type:decimal(10,2) unsigned;comment:'邮费'"`
	Weight      uint    `gorm:"comment:'重量：单位克'"` // 比如京东超过多少重量加收运费，数据库这里不用管这个逻辑，这个是程序的问题。这里只需要留出字段
}

// ToMap 实现 ToMap 方法
func (order *Order) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ID":          order.ID,
		"Code":        order.Code,
		"Type":        order.Type,
		"Amount":      order.Amount,
		"PaymentType": order.PaymentType,
		"Status":      order.Status,
	}
}

// RandomNum 生成随机数
func RandomNum(productId, userId uint) string {
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(productId))
	userNum := strconv.Itoa(int(userId))
	return number + productNum + userNum
}

// ShowOrder 根据订单号查询订单
func ShowOrder(code string) (Order, int) {
	var order Order
	err := db.DB.Where("code = ?", code).First(&order).Error
	if err != nil {
		logging.Info(err)
		return order, e.ERROR_DATABASE
	}
	return order, e.SUCCESS
}

// ListOrder 根据订单号查询订单
func ListOrder(id string) ([]Order, int) {
	var orders []Order
	err := db.DB.Where("user_id=?", id).Find(&orders).Error
	if err != nil {
		logging.Info(err)
		return orders, e.ERROR_DATABASE
	}
	return orders, e.SUCCESS
}

// ListenOrder 监听订单是否过期
func ListenOrder(ctx context.Context) {
	go func() {
		for {
			opt := redis.ZRangeBy{
				Min:    strconv.Itoa(0),
				Max:    strconv.Itoa(int(time.Now().Unix())),
				Offset: 0,
				Count:  10,
			}
			orderList, err := cache.RC.ZRangeByScore(ctx, os.Getenv("REDIS_ZSET_KEY"), &opt).Result()
			if err != nil {
				logging.Info("redis err:", err)
			}
			if len(orderList) != 0 {
				var numList []int
				for _, v := range orderList {
					i, err := strconv.Atoi(v)
					if err != nil {
						logging.Info("Atoi err:", err)
					}
					numList = append(numList, i)
				}
				if err := db.DB.Delete(&Order{}, "order_num IN (?)", numList).Error; err != nil {
					logging.Info("myql err:", err)
				}
				if err := cache.RC.ZRem(ctx, os.Getenv("REDIS_ZSET_KEY"), orderList).Err(); err != nil {
					logging.Info("redis err:", err)
				}
			}
		}
	}()
}
