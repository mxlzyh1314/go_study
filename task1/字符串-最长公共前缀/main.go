package main

import "fmt"

//编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ { // 遍历第一个字符串的每个字符
		char := strs[0][i] // 将获取当前字符赋值给char变量

		for j := 1; j < len(strs); j++ { // 遍历切片中第二个字符串的每个字符
			if i >= len(strs[j]) || strs[j][i] != char { // 判断i的长度是否大于等于j或者当前字符是否等于char
				return strs[0][:i] // 返回公共前缀
			}
		}
	}

	return strs[0]
}

func main() {

	s1 := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(s1)

	s2 := longestCommonPrefix([]string{"dog", "racecar", "car"})
	fmt.Println(s2)
}
