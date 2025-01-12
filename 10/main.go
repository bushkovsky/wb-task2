package main

import (
	"fmt"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
var tempature = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

func main() {

	// хранилище для разбиения по множествам
	groups := make(map[int][]float64)

	for _, el := range tempature {
		key := int(el) / 10 // для каждого ключа множества добавляем в соответсвующее нашу температуру
		groups[key*10] = append(groups[key*10], el)
	}

	for setKey, val := range groups {
		fmt.Print(setKey)
		fmt.Print(": ")
		for _, el := range val {
			fmt.Print(el)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
