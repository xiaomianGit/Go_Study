/**
* 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
 */
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func (counter *Counter) increment() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.value++
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("value:%d", counter.value)
}
