package main

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
// 每个协程对计数器进行1000次递增操作，最后输出计数器的值。

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 使用原子操作的计数器变量
	var counter int64
	var wg sync.WaitGroup

	// 启动10个协程
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 每个协程对计数器进行1000次递增操作
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终计数器值
	fmt.Printf("计数器最终值: %d\n", atomic.LoadInt64(&counter))
}
