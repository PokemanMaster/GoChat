package model

import (
	"IMProject/common/cache"
	"context"

	"time"
)

func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	cache.RC.Set(ctx, key, val, timeTTL)
}
