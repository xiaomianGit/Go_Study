/**
* 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
 */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func (counter *Counter) increment(wg *sync.WaitGroup) {
	counter.mutex.Lock()
	defer wg.Done()
	counter.value++
}

func main() {
	var counter uint32
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddUint32(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("value:%d", counter)
}
