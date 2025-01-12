package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	value int64
}

func main() {

	counter := Counter{value: 0}
	wg := sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(c Counter, n int) {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				atomic.AddInt64(&counter.value, 1)
				fmt.Println("+ 1", n)
			}

		}(counter, i)
	}

	wg.Wait()

	fmt.Println(counter.value)

}
