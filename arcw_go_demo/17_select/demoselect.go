package main

import (
	"fmt"
	"time"
)

/**
通道选择
select 与 channel结合起来使用
*/
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "给通道1的数据"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "给通道2的数据"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println("通道1接受：", m1)
		case m2 := <-c2:
			fmt.Println("通道2接受：", m2)

		}
	}
}
