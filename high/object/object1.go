/**
* 指针，修改引用地址值
 */
package main

import (
	"fmt"
)

func main() {
	rectangle := Rectangle{high: 32}
	circle := Circle{high: 3.2}
	fmt.Printf("rectangle:%.2f \n", rectangle.Area())
	fmt.Printf("rectangle:%.2f \n", rectangle.AAA())
	fmt.Printf("circle:%.2f \n", circle.Perimeter())

}

type Shape interface {
	Areae() float32
	Perimeter() float32
}

type Rectangle struct {
	high float32
}

type Circle struct {
	high float32
}

func (s *Rectangle) Area() float32 {
	return s.high
}
func (s *Rectangle) AAA() float32 {
	return s.high * 2
}
func (s *Circle) Perimeter() float32 {
	return s.high
}
