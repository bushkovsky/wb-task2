package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

// arr - исходный массив
var arr = [5]int{2, 4, 5, 8, 10}

func sum(in chan int, out chan int) {
	summ := 0
	for v := range in {
		summ += v
	}
	out <- summ
}

func main() {

	in := make(chan int)
	out := make(chan int)
	go sum(in, out)

	var wg sync.WaitGroup
	wg.Add(len(arr))

	for _, elem := range arr {
		go func(el int, in chan int) {
			defer wg.Done()
			in <- el * el
		}(elem, in)
	}

	wg.Wait()
	close(in)
	fmt.Println(<-out)
	close(out)
}
