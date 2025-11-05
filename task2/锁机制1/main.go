package main

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
// 每个协程对计数器进行1000次递增操作，最后输出计数器的值。

import (
	"fmt"
	"sync"
)

// Counter 计数器结构体
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment 递增计数器值
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// GetValue 获取计数器当前值
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// 创建共享计数器
	counter := &Counter{}
	var wg sync.WaitGroup
	
	// 启动10个协程
	wg.Add(10)
	
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 每个协程对计数器进行1000次递增操作
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	
	// 等待所有协程完成
	wg.Wait()
	
	// 输出最终计数器值
	fmt.Printf("计数器最终值: %d\n", counter.GetValue())
}