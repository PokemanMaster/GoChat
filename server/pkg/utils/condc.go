package utils

import (
	"fmt"
	"sync"
	"time"
)

var (
	buffer  []int // 缓冲区
	maxSize = 10  // 缓冲区最大容量
	lock    = sync.Mutex{}
	cond    = sync.NewCond(&lock)
)

func producer() {
	for i := 0; i < 20; i++ {
		lock.Lock()
		for len(buffer) == maxSize {
			cond.Wait() // 缓冲区满时等待
		}
		buffer = append(buffer, i)
		fmt.Printf("Produced: %d\n", i)
		cond.Signal() // 通知消费者
		lock.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(id int) {
	for {
		lock.Lock()
		for len(buffer) == 0 {
			cond.Wait() // 缓冲区为空时等待
		}
		item := buffer[0]
		buffer = buffer[1:]
		fmt.Printf("Consumer %d consumed: %d\n", id, item)
		cond.Signal() // 通知生产者
		lock.Unlock()
		time.Sleep(200 * time.Millisecond)
	}
}

func main_test() {
	go producer()
	go consumer(1)
	go consumer(2)

	// 主协程阻塞，防止退出
	select {}
}
