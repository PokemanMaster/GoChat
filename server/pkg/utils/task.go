package utils

import (
	"time"
)

type TimerFunc func(interface{}) bool

func Timer(delay, tick time.Duration, fun TimerFunc, param interface{}) {
	go func() {
		if fun == nil {
			return
		}
		// 创建了一个定时器，等待 delay 时间后，开始执行回调函数。
		t := time.NewTimer(delay)

		for {
			select {
			case <-t.C:
				// 当定时器触发时，检查 fun(param) 的返回值。
				// 如果返回 false，则停止定时器并退出。
				// 否则，继续重置定时器，并在 tick 时间间隔后再次执行。
				if fun(param) == false {
					return
				}
				// 通过 t.Reset(tick) 重置定时器，使它继续在 tick 间隔后触发。
				t.Reset(tick)
			}
		}
	}()
}
