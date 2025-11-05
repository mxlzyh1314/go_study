package main

import "fmt"

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

// Person 结构体包含姓名和年龄字段
type Person struct {
	Name string
	Age  int
}

// Employee 结构体通过组合 Person 结构体并添加 EmployeeID 字段
type Employee struct {
	Person     // 匿名字段组合 Person 结构体
	EmployeeID int
}

// PrintInfo 方法输出员工的完整信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n")
	fmt.Printf("姓名: %s\n", e.Name)
	fmt.Printf("年龄: %d\n", e.Age)
	fmt.Printf("员工ID: %d\n", e.EmployeeID)
}

func main() {
	// 创建 Employee 实例
	employee := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: 1001,
	}

	// 调用 PrintInfo 方法输出员工信息
	employee.PrintInfo()

}
