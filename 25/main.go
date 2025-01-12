package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func customSleep(dur time.Duration) {
	timer := time.NewTimer(dur)
	<-timer.C
}

func printTime(sec int) {
	for sec != 0 {
		fmt.Println(sec)
		time.Sleep(time.Second)
		sec--
	}
}

func main() {
	go printTime(5)
	customSleep(time.Second * 5)
}
