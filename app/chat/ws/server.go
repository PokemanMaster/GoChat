package ws

import (
	"IMProject/common/cache"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// 后端调度逻辑处理
func dispatch(data []byte) {
	// 反序列化
	msg := Message{}
	msg.CreateTime = uint64(time.Now().Unix())
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch msg.Type {
	case 1: // 好友发送的消息内容
		fmt.Println("dispatch  data :", string(data))
		sendMsg(msg.TargetId, data)
	case 2: // 群发送的消息内容
		sendGroupMsg(msg.TargetId, data)
	}
}

// 发送私信
func sendMsg(userId int64, data []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()

	// 反序列化
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 检查用户是否在线
	ctx := context.Background()
	targetIdStr := strconv.Itoa(int(userId))
	userIdStr := strconv.Itoa(int(msg.UserId))
	msg.CreateTime = uint64(time.Now().Unix())
	online, err := cache.Red.Get(ctx, "online_"+userIdStr).Result() // 获取用户的在线状态
	if err != nil {
		fmt.Println(err.Error())
	}

	// 发送消息到在线用户:
	if online != "" {
		if ok { // 如果用户在线并且在 clientMap 中找到对应的 Node，则将消息通过 node.DataQueue 发送给用户。
			node.DataQueue <- data
		}
	}

	// 构建消息的存储顺序键
	var key string
	if userId > msg.UserId {
		key = "msg_" + userIdStr + "_" + targetIdStr
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}

	// 然后将消息添加到 Redis 的有序集合中，以便以后检索历史消息
	res, err := cache.Red.ZRevRange(ctx, key, 0, -1).Result() // 获取当前对话中的历史消息
	if err != nil {
		fmt.Println(err.Error())
	}
	score := float64(cap(res)) + 1
	response, err := cache.Red.ZAdd(ctx, key, &redis.Z{score, data}).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(response)
}

// 发送群消息
func sendGroupMsg(targetId int64, msg []byte) {
	sendMsg(targetId, msg)
}
