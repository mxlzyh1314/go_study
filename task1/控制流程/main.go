package main

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，
//其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
//结合 if 条件判断和 map 数据结构来解决，
//例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

import "fmt"

func singleNumber(nums []int) int {   // 创建一个名为singleNumbe的函数，输入变量nums为整型切片，输出为整型
	countMap := make(map[int]int)   // 创建一个名为countMap的map，map的k和v均为整型，用于记录每个元素出现的次数
	for _, num := range nums {   // 遍历nums切片
		countMap[num]++   // 等同于countMap[num]=countMap[num]+1，将nums切片中的元素作为key，出现的次数作为value，存入countMap中
	}
	for num, count := range countMap {   // 遍历countMap
		if count == 1 {    // 判断value是否为1，即只出现一次的元素
			return num	  // 返回只出现一次的元素
		}
	}
	return 0   // 如果没有只出现一次的元素，则返回0
}

func main() {

	test1 := singleNumber([]int{2, 2, 1})
	fmt.Printf("此次出现一次的元素为：%d\n", test1)

	test2 := singleNumber([]int{4, 1, 2, 1, 2})
	fmt.Printf("此次出现一次的元素为：%d\n", test2)

	test3 := singleNumber([]int{1})
	fmt.Printf("此次出现一次的元素为：%d\n", test3)

}
