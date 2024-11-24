package common

import (
	"github.com/PokemanMaster/GoChat/app/chat/ws"
	MGroupBasic "github.com/PokemanMaster/GoChat/app/group/model"
	MUserBasic "github.com/PokemanMaster/GoChat/app/user/model"
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
			&MUserBasic.UserBasic{},
			&MUserBasic.Contact{},
			&MGroupBasic.GroupBasic{},
		)
	if err != nil {
		fmt.Println("err", err)
	}
}
