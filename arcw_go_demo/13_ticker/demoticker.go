package main

import (
	"fmt"
	"time"
)

/**
timer: sleep 等待并执行
ticker:轮训执行
*/

func main() {
	ticker := time.NewTicker(time.Millisecond * 1000)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t.Format("2006-01-02 15:04:05.99999"))
		}
	}()

	// 等待一会并停止ticker
	time.Sleep(time.Second * 6)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
