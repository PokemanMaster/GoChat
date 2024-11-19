package ws

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func (msg Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(msg)
}

// Heartbeat 更新用户心跳
func (node *Node) Heartbeat(currentTime uint64) {
	node.HeartbeatTime = currentTime
	return
}

// IsHeartbeatTimeOut 用户心跳是否超时
func (node *Node) IsHeartbeatTimeOut(currentTime uint64) (timeout bool) {
	if node.HeartbeatTime+viper.GetUint64("timeout.HeartbeatMaxTime") <= currentTime {
		fmt.Println("心跳超时自动下线", node)
		timeout = true
	}
	return
}

// CleanConnection 清理超时连接
func CleanConnection(param interface{}) (result bool) {
	result = true
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("cleanConnection err", err)
		}
	}()
	currentTime := uint64(time.Now().Unix())
	for i := range clientMap {
		node := clientMap[i]
		if node.IsHeartbeatTimeOut(currentTime) {
			fmt.Println("心跳超时关闭连接：", node)
			node.Conn.Close()
		}
	}
	return result
}
