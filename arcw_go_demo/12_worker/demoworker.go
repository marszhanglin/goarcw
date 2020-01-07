package main

import (
	"fmt"
	"time"
)

/**
worker
*/

// 我们将在worker函数里面运行几个实例，这个函数从jobs通道里面接受任务，
// 然后把运行结果发送到result通道。 每个job我们都休眠一会儿，来模拟一个耗时任务
func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// 定义两个通道
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动3个worker协程，一开始的时候wirker阻塞执行，
	// 因为jobs通道里面还没有工作任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送9个任务，然后关闭通道，告知任务发送完成
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}

}
