package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/
var ch1 = make(chan int)
var ch2 = make(chan int)
var arr = [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

func sq(in, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printNum(in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go sq(ch1, ch2, &wg)
	go printNum(ch2, &wg)

	for i := 0; i < len(arr); i++ {
		ch1 <- arr[i]
	}

	close(ch1)
	wg.Wait()
}
