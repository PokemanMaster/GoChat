package main

import (
	"IMProject/common"
	"IMProject/router"
<<<<<<< HEAD
	"fmt"
)

func main() {
	common.Init() // 初始化配置
	fmt.Println("123121221233")
=======
)

func main() {
	common.Init()         // 初始化配置
>>>>>>> 654d6dd65ef53e6a8b34a0b7d48c3276960bbc4c
	r := router.Router()  // 初始化路由
	panic(r.Run(":9000")) // 初始化监听端口
}
