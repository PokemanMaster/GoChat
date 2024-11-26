package service

import (
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
)

type GetMessageService struct {
	UserIdA int
	UserIdB int
	Start   int
	End     int
	IsRev   bool
}

// Get 获取消息记录
func (service *GetMessageService) Get() *resp.Response {
	userIdA := service.UserIdA
	userIdB := service.UserIdB
	start := service.Start
	end := service.End
	isRev := service.IsRev
	// 通过调用 models.RedisMsg 获取 Redis 中的消息记录
	res := GetMessage(int64(userIdA), int64(userIdB), int64(start), int64(end), isRev)
	// 将查询到的消息结果返回给客户端
	return &resp.Response{
		Status: e2.SUCCESS,
		Msg:    e2.GetMsg(e2.SUCCESS),
		Data:   res,
	}
}
