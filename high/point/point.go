/**
* 指针，修改引用地址值
 */
package main

import "fmt"

func main() {
	var num = 5
	receipt(&num)
	fmt.Println(num)
}

func receipt(num *int) {
	*num = *num + 10
}
