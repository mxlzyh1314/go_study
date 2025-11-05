package main

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个无缓冲通道用于传输整数
	ch := make(chan int)
	var wg sync.WaitGroup
	
	// 添加两个等待组计数
	wg.Add(2)
	
	// 协程1：生成数字并发送到通道
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i // 将数字发送到通道
		}
		close(ch) // 发送完所有数字后关闭通道
	}()
	
	// 协程2：从通道接收数字并打印
	go func() {
		defer wg.Done()
		for num := range ch { // 从通道接收数字直到通道关闭
			fmt.Printf("接收到数字: %d\n", num)
		}
	}()
	
	// 等待两个协程完成
	wg.Wait()
	fmt.Println("所有数字处理完成")
}