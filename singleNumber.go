package main

import "fmt"

//136. 只出现一次的数字
func main() {
	fmt.Print(singleNumber([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 7, 8, 8}))
}

func singleNumber(nums []int) int {
	queue := 0
	for _, ele := range nums {
		queue ^= ele
	}
	return queue
}
