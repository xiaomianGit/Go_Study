/**
*指针，修改切片引用地址值
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	receipt(&nums)
	fmt.Println(nums)
}

func receipt(nums *[]int) {
	for index, _ := range *nums {
		(*nums)[index] = (*nums)[index] * 2
	}
}
