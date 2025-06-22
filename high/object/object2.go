/**
* 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
 */
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  uint32
}

type Employee struct {
	persons    []Person
	EmployeeID uint
}

func (Employee Employee) PrintInfo() {
	for _, Person := range Employee.persons {
		fmt.Printf("姓名：%s,年龄:%d \n", Person.Name, Person.Age)

	}
}

func main() {
	Employee := Employee{
		EmployeeID: 123,
		persons:    []Person{Person{Name: "张三", Age: 22}, Person{Name: "kavin", Age: 23}},
	}
	Employee.PrintInfo()

}
