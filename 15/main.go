package main

import "fmt"

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.
*/

/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}
func main() {
  someFunc()
}

________________
Объяснение:
1) может произойти утечка памяти, так как не происходит копирования части строки v в
justString, а просто justString ссылается на первые 100 байт, таким образом происходит
удержание ссылки на всю строку v и GC ее не подчищает

Решение:
Надо копировать строку
*/

var justString string

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100]) // копирование
	fmt.Println(justString)
}

func main() {
	someFunc()
}
