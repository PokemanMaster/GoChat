package resp

import (
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"log"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataTotal 带有总数的Data结构
type DataTotal struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// DataToken 带有token的Data结构
type DataToken struct {
	Data  interface{} `json:"data"`
	Token string      `json:"token"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// BuildResponseTotal 带有总数的列表构建器
func BuildResponseTotal(items interface{}, total uint) Response {
	return Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data: DataTotal{
			Items: items,
			Total: total,
		},
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
