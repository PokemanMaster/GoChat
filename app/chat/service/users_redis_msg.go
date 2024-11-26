package service

import (
	"context"
	"fmt"
	"github.com/PokemanMaster/GoChat/server/common/cache"
	"strconv"
	"sync"
)

var rwLocker sync.RWMutex // 读写锁

// GetMessage 获取缓存里面的消息
func GetMessage(userIdA int64, userIdB int64, start int64, end int64, isRev bool) []string {
	rwLocker.RLock()
	rwLocker.RUnlock()
	ctx := context.Background()
	userIdStr := strconv.Itoa(int(userIdA))
	targetIdStr := strconv.Itoa(int(userIdB))

	var key string
	if userIdA > userIdB {
		key = "msg_" + targetIdStr + "_" + userIdStr
	} else {
		key = "msg_" + userIdStr + "_" + targetIdStr
	}

	var rels []string
	var err error
	if isRev { // key 从低到高
		rels, err = cache.RC.ZRange(ctx, key, start, end).Result()
	} else { // key 从高到低
		rels, err = cache.RC.ZRevRange(ctx, key, start, end).Result()
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	return rels
}
