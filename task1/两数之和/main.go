package main

// 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
import "fmt"

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	s1 := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(s1)

	s2 := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(s2)

	s3 := twoSum([]int{3, 3}, 6)
	fmt.Println(s3)
}
