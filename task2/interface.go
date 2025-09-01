package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	Width  float32
	Height float32
}

func (r *Rectangle) Area() float32 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float32 {
	return (r.Width + r.Height) * 2
}

type Circle struct {
	Radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

func main() {
	//创建 Rectangle 实例
	rectangle := Rectangle{
		Width:  5.0,
		Height: 2.0,
	}
	//创建 Circle 实例
	circle := Circle{
		Radius: 4.0,
	}

	// 使用接口变量
	var shapes []Shape
	shapes = append(shapes, &rectangle, &circle)

	// 遍历调用接口方法
	fmt.Println("=== 通过接口调用 ===")
	for i, shape := range shapes {
		fmt.Printf("Shape %d:\n", i+1)
		fmt.Printf("  Area: %.2f\n", shape.Area())
		fmt.Printf("  Perimeter: %.2f\n", shape.Perimeter())
		fmt.Println()
	}

	// 直接调用具体类型的方法
	fmt.Println("=== 直接调用 ===")
	fmt.Printf("Rectangle - Width: %.2f, Height: %.2f\n", rectangle.Width, rectangle.Height)
	fmt.Printf("  Area: %.2f\n", rectangle.Area())
	fmt.Printf("  Perimeter: %.2f\n", rectangle.Perimeter())

	fmt.Printf("\nCircle - Radius: %.2f\n", circle.Radius)
	fmt.Printf("  Area: %.2f\n", circle.Area())
	fmt.Printf("  Perimeter: %.2f\n", circle.Perimeter())
}
