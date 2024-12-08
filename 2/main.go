package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых
из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

// arr - исходный массив
var arr = [5]int{2, 4, 5, 8, 10}

func main() {

	// wait group для возможности дождаться завершения выполнения всех горутин основной main()
	var wg sync.WaitGroup
	wg.Add(len(arr))

	for _, elem := range arr {
		// запускаем воркер, который рассчитывает квадрат elem
		go func(el int) {

			defer wg.Done()
			fmt.Println(el, el*el)

		}(elem)
	}

	// ждем завершения всех горутин
	wg.Wait()
}
