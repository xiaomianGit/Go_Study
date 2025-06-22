/**
* 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
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
	for i := 1; i <= 100; i++ {
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
