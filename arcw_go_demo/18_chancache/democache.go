package main

import "fmt"

/**
通道缓冲
*/
func main() {
	// 通道缓冲区的大小为2
	c := make(chan string, 4)

	// 发送三个数据
	c <- "data1"
	c <- "data2"
	c <- "data3"
	c <- "data4"

	// 获取通道数据
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
