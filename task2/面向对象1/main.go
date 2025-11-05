package main

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

import (
	"fmt"
	"math"
)

// Shape 定义形状接口，包含面积和周长计算方法
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}
// Area 计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}



// Area 计算圆形面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆形周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建矩形实例
	rectangle := Rectangle{Width: 5.0, Height: 3.0}
	
	// 创建圆形实例
	circle := Circle{Radius: 4.0}
	
	// 通过接口使用多态特性
	shapes := []Shape{rectangle, circle}
	fmt.Printf("\n通过接口调用:\n")
	for _, shape := range shapes {
		fmt.Printf("面积: %.2f, 周长: %.2f\n", shape.Area(), shape.Perimeter())
	}
}