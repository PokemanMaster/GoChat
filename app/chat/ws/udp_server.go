package ws

import (
	"fmt"
	"github.com/spf13/viper"
	"net"
)

func init() {
	go udpSendProc()
	go udpRecvProc()
	fmt.Println("udp Init")
}

var udpsendChan = make(chan []byte, 1024) // 定义 udp 通道

// 把数据发送到udp通道中
func broadMsg(data []byte) {
	udpsendChan <- data
}

// udpSendProc UDP发送数据
func udpSendProc() {
	// 建立一个 UDP 连接
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: viper.GetInt("port.udp"),
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 数据发送逻辑
	for {
		select {
		case data := <-udpsendChan:
			_, err = con.Write(data) // 发送 UDP 数据包
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}

// udpRecvProc UDP接收数据
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: viper.GetInt("port.udp"),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	defer con.Close()

	// 数据接收逻辑
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:]) // 读取数据到缓冲区
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}
