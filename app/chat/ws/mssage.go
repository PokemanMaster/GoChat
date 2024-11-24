package ws

import (
	"encoding/json"
	"fmt"
	Muser "github.com/PokemanMaster/GoChat/app/user/model"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	clientMap = make(map[int64]*Node) // 用户与节点映射
	rwLocker  sync.RWMutex            // 读写锁

	nodePool = sync.Pool{ // Node 对象池
		New: func() interface{} {
			return &Node{
				DataQueue: make(chan []byte, 50),
				GroupSets: set.New(set.ThreadSafe),
			}
		},
	}

	messagePool = sync.Pool{ // Message 对象池
		New: func() interface{} {
			return &Message{}
		},
	}
)

// Message 用户发送数据的格式
type Message struct {
	gorm.Model
	UserId     int64  // 发送者
	TargetId   int64  // 接受者
	Type       int    // 发送类型  1私聊  2群聊  3心跳
	Media      int    // 消息类型  1文字  2表情包 3语音  4图片/表情包
	Content    string // 消息内容
	CreateTime uint64 // 创建时间
	ReadTime   uint64 // 读取时间
	Pic        string // 存储图片消息的地址
	Url        string // 存储消息中附带的超链接地址
	Desc       string // 为消息提供额外的描述或摘要信息
	Amount     int    // 表示消息中的某些数字值或计数信息
}

// Node 存储用户节点信息
type Node struct {
	Conn          *websocket.Conn //连接
	Addr          string          //客户端地址
	FirstTime     uint64          //首次连接时间
	HeartbeatTime uint64          //心跳时间
	LoginTime     uint64          //登录时间
	DataQueue     chan []byte     //消息
	GroupSets     set.Interface   //好友 / 群
}

// Chat 处理 WebSocket 连接
func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := query.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)

	isValid := true
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return isValid },
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建 Node 对象
	currentTime := uint64(time.Now().Unix())
	node := nodePool.Get().(*Node)
	node.Conn = conn
	node.Addr = conn.RemoteAddr().String()
	node.HeartbeatTime = currentTime
	node.LoginTime = currentTime

	// 加入全局映射
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	// 启动消息处理逻辑
	go sendProc(node)
	go recvProc(node)

	// 缓存用户在线信息
	Muser.SetUserOnlineInfo("online_"+id, []byte(node.Addr), time.Duration(viper.GetInt("timeout.RedisOnlineTime"))*time.Hour)
}

// sendProc 处理消息发送
func sendProc(node *Node) {
	defer releaseNode(node) // 确保资源释放
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// recvProc 处理消息接收
func recvProc(node *Node) {
	// 移除用户映射并释放资源
	defer func() {
		rwLocker.Lock()
		for id, n := range clientMap {
			if n == node {
				delete(clientMap, id)
			}
		}
		rwLocker.Unlock()
		releaseNode(node)
	}()

	// 用于存储分片消息的缓存
	var messageCache = make(map[string][]byte)

	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 从消息中提取分片标识符
		msg := messagePool.Get().(*Message)
		err = json.Unmarshal(data, msg)
		if err != nil {
			messagePool.Put(msg)
			continue
		}

		// 心跳消息处理
		if msg.Type == 3 {
			node.HeartbeatTime = uint64(time.Now().Unix())
		} else if len(data) > maxMessageSize {
			// 如果消息大于 maxMessageSize，则进行分片处理
			// 假设消息的 Content 中有一个字段标识分片的唯一 ID 和当前分片的索引
			// 比如：`messageId` 用于标识分片消息的整体，`chunkIndex` 用于标识每个分片的序号
			// 这里我们将这个信息保存到缓存中，等所有分片接收完再处理

			// 使用消息中的 `messageId` 来区分不同的分片
			messageId := msg.Content // 假设 Content 字段是用于标识消息 ID
			if _, exists := messageCache[messageId]; !exists {
				messageCache[messageId] = []byte{} // 初始化缓存
			}
			// 将当前分片追加到缓存
			messageCache[messageId] = append(messageCache[messageId], data...)

			// 如果该消息已经接收完（假设可以通过 Content 的长度来判断是否是最后一个分片）
			// 这里假设每个分片都有一个长度标识，我们用 `msg.Amount` 来表示
			if len(messageCache[messageId]) == msg.Amount {
				// 合并所有分片，进行处理
				completeMessage := messageCache[messageId]
				dispatch(completeMessage)       // 将合并后的完整消息分发处理
				delete(messageCache, messageId) // 清理缓存
			}
		} else {
			// 不是分片消息，直接分发处理
			dispatch(data)
		}

		messagePool.Put(msg)
	}
}

// releaseNode 释放 Node 对象
func releaseNode(node *Node) {
	node.Conn = nil
	node.GroupSets.Clear()
	close(node.DataQueue)
	nodePool.Put(node)
}
