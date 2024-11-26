package main

import (
	"github.com/PokemanMaster/GoChat/server/common"
	"github.com/PokemanMaster/GoChat/server/router"
)

func main() {
	common.Init()         // 初始化配置
	r := router.Router()  // 初始化路由
	panic(r.Run(":9000")) // 初始化监听端口
}
