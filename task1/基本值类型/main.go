package main

import (
	"fmt"
)

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
//
//将大整数加 1，并返回结果的数字数组。

func plusOne(digits []int) []int {

	// 从最后一位开始加1
	for i := len(digits) - 1; i >= 0; i-- { // 从最后一位开始遍历
		if digits[i] < 9 { // 判断当前位是否小于9,避免出现9+1进位的情况
			digits[i]++   // 当前位加1
			return digits // 返回结果
		}
		digits[i] = 0 // 如果不满足if判断，说明当前位已经等于9，则将当前位置为0
	}

	// 如果所有位都是9，需要扩展数组
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

func main() {

	s1 := plusOne([]int{1, 2, 3})
	fmt.Println(s1)

	s2 := plusOne([]int{4, 3, 2, 1})
	fmt.Println(s2)

	s3 := plusOne([]int{9})
	fmt.Println(s3)
}
