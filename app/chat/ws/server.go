package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

const maxMessageSize = 1024 // 设置每个消息的最大大小为 1KB

// dispatch 消息调度
func dispatch(data []byte) {
	msg := messagePool.Get().(*Message)
	defer messagePool.Put(msg)

	err := json.Unmarshal(data, msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch msg.Type {
	case 1: // 私信
		sendSingleChunk(msg.TargetId, data)
	case 2: // 群聊
		sendGroupMsg(msg.TargetId, data)
	}
}

// 分片发送消息
func sendSingleChunk(userId int64, data []byte) {
	if len(data) > maxMessageSize {
		// 如果消息过大，进行分片发送
		for i := 0; i < len(data); i += maxMessageSize {
			end := i + maxMessageSize
			if end > len(data) {
				end = len(data)
			}
			chunk := data[i:end]
			sendMsg(userId, chunk)
		}
	} else {
		// 如果消息不大，直接发送
		sendMsg(userId, data)
	}
}

// 发送私信
func sendMsg(userId int64, data []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()

	// 反序列化消息
	msg := messagePool.Get().(*Message) // 使用对象池
	defer messagePool.Put(msg)          // 处理完成后放回池中

	err := json.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("Message Unmarshal Error:", err)
		return
	}

	// 检查用户在线状态
	ctx := context.Background()
	targetIdStr := strconv.Itoa(int(userId))
	userIdStr := strconv.Itoa(int(msg.UserId))
	msg.CreateTime = uint64(time.Now().Unix())
	online, err := cache.RC.Get(ctx, "online_"+userIdStr).Result()
	if err != nil {
		fmt.Println("Redis Online Check Error:", err)
	}

	// 如果用户在线，发送消息到 WebSocket
	if online != "" && ok {
		node.DataQueue <- data
	}

	// 构建 Redis 的存储键
	var key string
	if userId > msg.UserId {
		key = "msg_" + userIdStr + "_" + targetIdStr
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}

	// 存储消息到 Redis 的有序集合
	score := float64(time.Now().UnixNano())
	_, err = cache.RC.ZAdd(ctx, key, &redis.Z{Score: score, Member: data}).Result()
	if err != nil {
		fmt.Println("Redis ZAdd Error:", err)
	}
}

// 发送群消息
func sendGroupMsg(targetId int64, msg []byte) {
	sendMsg(targetId, msg)
}
