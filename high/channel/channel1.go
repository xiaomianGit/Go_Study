/**
* 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
 */
package main

import (
	"fmt"
	"sync"
)

func main() {
	chan1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go generatorNum(chan1, &wg)
	go readNum(chan1, &wg)
	wg.Wait()
}

func generatorNum(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成时通知WaitGroup
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c) // 发送完毕后关闭通道
}
func readNum(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成时通知WaitGroup
	for num := range c {
		fmt.Printf("num:%d \n", num)
	}
}
