package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

var str = "snow dog sun"

func main() {

	words := strings.Split(str, " ")
	strTmp := strings.Builder{}
	for i := 0; i < len(words); i++ {
		strTmp.WriteString(words[len(words)-i-1])
		strTmp.WriteString(" ")
	}

	fmt.Println(strTmp.String())

}
