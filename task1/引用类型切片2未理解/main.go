package main

//56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
//请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，
//遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
//如果没有重叠，则将当前区间添加到切片中。

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 处理空数组情况
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果数组
	merged := [][]int{}

	// 遍历所有区间
	for _, interval := range intervals {
		// 如果结果数组为空，或者当前区间与最后一个区间不重叠
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			// 合并区间：更新最后一个区间的结束位置
			if interval[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = interval[1]
			}
		}
	}

	return merged
}

func main() {
	// 测试示例1
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("示例1输入:", intervals1)
	fmt.Println("示例1输出:", merge(intervals1))

	// 测试示例2
	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println("示例2输入:", intervals2)
	fmt.Println("示例2输出:", merge(intervals2))

	// 测试示例3
	intervals3 := [][]int{{4, 7}, {1, 4}}
	fmt.Println("示例3输入:", intervals3)
	fmt.Println("示例3输出:", merge(intervals3))
}
