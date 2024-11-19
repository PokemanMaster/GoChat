package main

import (
	"IMProject/common"
	"IMProject/router"
	"fmt"
)

func main() {
	common.Init()           // 初始化配置
	fmt.Println("12312123") // 打印测试信息
	r := router.Router()    // 初始化路由
	panic(r.Run(":9000"))   // 初始化监听端口
}
