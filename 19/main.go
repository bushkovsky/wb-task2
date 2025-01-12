package main

import "fmt"

func reverse(str string) (reverseStr string) {
	tmpStr := []rune{}

	strRunes := []rune(str)

	for i := len(strRunes) - 1; i > -1; i-- {

		tmpStr = append(tmpStr, strRunes[i])
	}

	reverseStr = string(tmpStr)
	return
}

func main() {
	fmt.Println(reverse("Hello"))
}
