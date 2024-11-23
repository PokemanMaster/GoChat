package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"strconv"
	"sync"
)

var RC *redis.Client
var once sync.Once // 避免了重复创建连接

const (
	// RankKey 每日排行
	RankKey = "rank"
	// ElectricalRank 家电排行
	ElectricalRank = "elecRank"
	// AccessoryRank 配件排行
	AccessoryRank = "acceRank"
	PublishKey    = "websocket"
)

// InitRedis 初始化 Redis
func InitRedis() {
	once.Do(func() {
		RC = redis.NewClient(&redis.Options{
			Addr:         viper.GetString("redis.addr"),
			Password:     viper.GetString("redis.password"),
			DB:           viper.GetInt("redis.DB"),
			PoolSize:     viper.GetInt("redis.poolSize"),
			MinIdleConns: viper.GetInt("redis.minIdleConn"),
		})
		pong, err := RC.Ping(context.Background()).Result()
		if err != nil {
			fmt.Println("Init Redis err", err)
		} else {
			fmt.Println("Init Redis", pong)
		}
	})
}

// ProductViewKey 商品点击数的key
func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	err := RC.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := RC.Subscribe(ctx, channel)
	fmt.Println("Subscribe", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe", msg.Payload)
	return msg.Payload, err
}
