package main

import (
	"github.com/PokemanMaster/GoChat/v1/server/common"
	"github.com/PokemanMaster/GoChat/v1/server/router"
)

func main() {
	common.Init()        // 初始化配置
	r := router.Router() // 初始化路由

	//certFile := "/common/nginx/lvyouwang.xyz.pem"
	//keyFile := "/common/nginx/lvyouwang.xyz.key"
	//if err := r.RunTLS(":9000", certFile, keyFile); err != nil {
	//	panic(err)
	//}

	panic(r.Run(":9000")) // 初始化监听端口
}
