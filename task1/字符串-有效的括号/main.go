package main

import "fmt"

/*
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

func isValid(s string) bool {
	// 创建栈存储左括号
	stack := []byte{}

	// 定义括号映射关系
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 遍历字符串中的每个字符
	for i := 0; i < len(s); i++ {
		char := s[i]

		// 如果是右括号
		if pair, exists := pairs[char]; exists {
			// 检查栈是否为空或者不匹配
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 是左括号，压入栈中
			stack = append(stack, char)
		}
	}

	// 最后栈应该为空才表示所有括号都正确匹配
	return len(stack) == 0
}

func main() {
	s1 := isValid("()")
	fmt.Println(s1)

	s2 := isValid("()[]{}")
	fmt.Println(s2)

	s3 := isValid("(]")
	fmt.Println(s3)

	s4 := isValid("([])")
	fmt.Println(s4)

	s5 := isValid("([)]")
	fmt.Println(s5)

}
