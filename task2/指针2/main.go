package main

import "fmt"

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。

func multiplyByTwo(slice *[]int) {
	// 遍历切片中的每个元素并乘以2
	for i := range *slice {
		(*slice)[i] *= 2
	}
}

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	// 输出原始切片
	fmt.Println("原始切片:", numbers)

	// 调用multiplyByTwo函数，传入切片的指针
	multiplyByTwo(&numbers)

	// 输出修改后的切片
	fmt.Println("修改后切片:", numbers)
}
