package common

import (
	"IMProject/app/chat/ws"
	MGroupBasic "IMProject/app/group/model"
	MUserBasic "IMProject/app/user/model"
	"IMProject/common/cache"
	"IMProject/common/db"
	"IMProject/config"
	"IMProject/pkg/utils"
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
