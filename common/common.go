package common

import (
	"github.com/PokemanMaster/GoChat/app/carousel/model"
	cart "github.com/PokemanMaster/GoChat/app/cart/model"
	model2 "github.com/PokemanMaster/GoChat/app/category/model"
	"github.com/PokemanMaster/GoChat/app/chat/ws"
	favorite "github.com/PokemanMaster/GoChat/app/favorite/model"
	MGroup "github.com/PokemanMaster/GoChat/app/group/model"
	order "github.com/PokemanMaster/GoChat/app/order/model"
	product "github.com/PokemanMaster/GoChat/app/product/model"
	transport "github.com/PokemanMaster/GoChat/app/transport/model"
	user "github.com/PokemanMaster/GoChat/app/user/model"
	"github.com/PokemanMaster/GoChat/common/logger"
	"go.uber.org/zap"

	warehouse "github.com/PokemanMaster/GoChat/app/warehouse/model"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/cache/rabbit"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/config"
	"github.com/PokemanMaster/GoChat/pkg/mid"
	"github.com/PokemanMaster/GoChat/pkg/utils"

	"fmt"
	"github.com/spf13/viper"
	"time"
)

func Init() {
	config.InitConfig() // 初始化 配置
	db.InitDB()         // 初始化 数据库
	migration()         // 初始化 数据表
	cache.InitRedis()   // 初始化 Redis
	//pb.UserPB()        // 初始化 用户 pb
	TimingCleanMysql() // 初始化 定时器

	// 初始化布隆过滤器
	mid.InitBloomFilter(10000, 3)
	for i := 1; i <= 50; i++ {
		itemID := fmt.Sprintf("%d", i)
		mid.BloomFilterGlobal.Add(itemID)
	}

	// 初始化日志
	logger.InitLogger(
		"./logs/app.log", // 日志文件路径
		10,               // 每个日志文件最大大小（MB）
		5,                // 最大保留旧日志文件数量
		30,               // 最大保留天数
		true,             // 是否压缩旧日志
		zap.DebugLevel,   // 日志级别
	)

	// 初始化rabbit消息队列
	rabbit.InitRabbitMQ()
	//go dao.Consumer1()
	//time.Sleep(2 * time.Second)
	//go dao.Consumer2()
}

// TimingCleanMysql 初始化定时器，定时清理数据库的超时连接
func TimingCleanMysql() {
	utils.Timer(
		time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second,
		time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second,
		ws.CleanConnection, "")
}

func migration() {
	err := db.DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&user.User{},
			&user.Contact{},
			&MGroup.GroupBasic{},
			&model.Carousel{},
			&cart.Cart{},
			&favorite.Favorite{},
			&order.Order{},
			&order.OrderDetail{},
			&product.Product{},              // 商品表
			&product.ProductBrand{},         // 商品品牌表
			&model2.ProductCategory{},       // 商品分类表
			&product.ProductCategoryBrand{}, // 商品分类与品牌关联表
			&product.ProductParam{},         // 商品参数表
			&transport.TransportBackstock{}, // 退货表
			&transport.TransportDelivery{},  // 快递表
			&user.User{},                    // 用户表
			&user.UserAddress{},             // 用户地址
			&user.UserLevel{},               // 用户等级
			&user.UserRating{},              // 用户评价表
			&warehouse.Warehouse{},          // 仓库表
			&warehouse.WarehouseProduct{},   // 仓库商品库存表
		)
	if err != nil {
		fmt.Println("err", err)
	}
}
