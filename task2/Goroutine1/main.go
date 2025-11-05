package main

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。

import (
	"fmt"
	"sync"
)

func odd(num int) {
	for i := 1; i <= num; i += 2 {
		fmt.Printf("奇数: %d\n", i)
	}
}

func even(num int) {
	for i := 2; i <= num; i += 2 {
		fmt.Printf("偶数: %d\n", i)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		odd(10)
	}()

	go func() {
		defer wg.Done()
		even(10)
	}()

	wg.Wait()
}
