/**
* 指针，修改引用地址值
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	go single()
	go double()
	time.Sleep(3 * time.Second)
}

func single() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println("奇数：", i)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
func double() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("偶数：", i)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
