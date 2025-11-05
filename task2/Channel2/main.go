package main

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
// 消费者协程从通道中接收这些整数并打印。// 

import (
	"fmt"
	"sync"
)


func main() {
	// 创建一个带缓冲的通道，缓冲区大小为10
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	
	// 添加生产者和消费者两个协程的等待计数
	wg.Add(2)
	
	// 生产者协程：向通道发送100个整数
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i // 发送整数到通道
			fmt.Printf("生产者发送: %d\n", i)
		}
		close(ch) // 发送完成后关闭通道
	}()
	
	// 消费者协程：从通道接收并打印整数
	go func() {
		defer wg.Done()
		for num := range ch { // 从通道接收直到通道关闭
			fmt.Printf("消费者接收: %d\n", num)
		}
	}()
	
	// 等待两个协程完成
	wg.Wait()
	fmt.Println("所有数据处理完成")
}