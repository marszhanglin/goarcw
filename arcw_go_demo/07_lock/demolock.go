package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/**
sync
*/
func main() {

	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var ops int64 = 0

	for r := 0; r < 100; r++ {

		go func() {
			total := 0
			for {
				// 对于每次读取，我们选取一个key来访问，mutex的lock函数用来保证对状态对唯一性访问，
				// 访问结束后，使用unlock来解锁，然后增加ops计数器
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				// 为保证这个协程不会让调度器处于饥饿状态，我们显示的使用runtime.Gosched释放了资源的控制权
				// 这种控制权会在通道结束或者time.sleep结束后释放，但是这里我们要手动地释放资源控制权
				runtime.Gosched()
			}

		}()

	}

	// 同样使用10个协程来模拟写状态
	for w := 0; w < 10; w++ {
		go func() {

			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}

		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

}
