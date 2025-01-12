package main

import (
	"fmt"
	"sync"
	"time"
)

var n int // количество секунд f
var wg sync.WaitGroup

// printSecond - печатает число, которое получает через канал
func printSecond(ch chan int) {
	defer wg.Done()

	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	n = 5
	ch := make(chan int)
	wg.Add(1)
	go printSecond(ch)

	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		ch <- i
	}

	close(ch)
	wg.Wait()

}
