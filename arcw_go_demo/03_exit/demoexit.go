package main

import (
	"fmt"
	"os"
)

/**
exit
使用os.Exit可以给定一个状态，然后立刻退出程序运行
*/
func main() {

	// 使用os.Exit的时候defer操作不会被执行
	defer fmt.Println("!")
	// 退出并设置状态值
	os.Exit(5)

}
