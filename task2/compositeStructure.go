// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

// 为 Employee 结构体实现一个 PrintInfo() 方法
func (e *Employee) PrintInfo() {
	fmt.Printf("Employee ID: %d\n", e.EmployeeID)
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
}
func main() {
	//创建 Employee 实例
	employee := Employee{
		Person: Person{
			Name: "jack",
			Age:  30,
		},
		EmployeeID: 100,
	}
	employee.PrintInfo()
	fmt.Println()

	// 直接访问组合的字段
	fmt.Println("直接访问字段:")
	fmt.Printf("姓名: %s\n", employee.Name) // 通过组合直接访问
	fmt.Printf("年龄: %d\n", employee.Age)  // 通过组合直接访问
	fmt.Printf("员工ID: %d\n", employee.EmployeeID)

	// 也可以通过 Person 字段访问
	fmt.Println("通过 Person 字段访问:")
	fmt.Printf("姓名: %s\n", employee.Person.Name)
	fmt.Printf("年龄: %d\n", employee.Person.Age)
}
