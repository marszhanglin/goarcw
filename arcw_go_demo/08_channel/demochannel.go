package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/**
channel

*/
func main() {

	// Go信号通知通过向一个channel发送os.Singnal来实现，我们将创建一个channel来接受这些通知，
	// 同时我们还用一个channel来在程序可以退出的时候通知我们
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify 在给定的channel上面注册该channel
	// 可以接受的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 这个gorountine阻塞等待信号的到来，当信号到来的时候
	// 输出该信号，然后通知程序可以结束了
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

// 主协程跟子协程同时阻塞，外部发送信号ctrl+c给子协程，子协程发送信号给主协程
