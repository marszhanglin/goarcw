package main

import (
	"fmt"
)

/**
关闭通道：该通道将不再允许写入数据。这个方法可以让通道数据的接受端知道数据已经全部发送完成了
*/
func main() {

	// 数据发送完成后，关闭jobs通道
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // 当前信号 是否有更多信号
			if more {
				fmt.Println("receive job", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true // 让主协程退出
				return
			}
		}
	}()

	// 这里向jobs通道写入三个数据，然后关闭通道
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done // 主协程等待
}
