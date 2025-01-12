package main

import "fmt"

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

var uniqe = make(map[rune]bool)

func findUniqe(str string) bool {
	tmp := []rune(str)
	for i := 0; i < len(tmp); i++ {
		if uniqe[tmp[i]] {
			return false
		} else {
			uniqe[tmp[i]] = true
		}
	}
	return true
}

func main() {
	fmt.Println(findUniqe("abcd"))
	fmt.Println(findUniqe("abCdefAaf"))
	fmt.Println(findUniqe("aabcd"))
}
