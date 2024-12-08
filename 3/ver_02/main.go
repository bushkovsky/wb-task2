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

func main() {

	sumChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(arr)) // Добавление в счетчик длины массива

	for _, element := range arr {
		go func(el int) {
			defer wg.Done()
			sumChan <- el * el // передача основной горутине квадрата числа
		}(element)
	}

	// запускаем горутину для ожидания завершения всех горутин считающих квадраты
	go func() {
		wg.Wait()
		close(sumChan)
	}()

	sum := 0
	// пока канал не закрыт (т.е. пока запущена хотя бы одна горутина, считающая квадрат числа)
	// достаем квадраты из канала и считаем сумму
	for num := range sumChan {
		sum += num
	}

	//канал закрыт, вышли из цикла, поэтому выводим сумму в stdout
	fmt.Println(sum)
}
