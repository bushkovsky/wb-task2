package main

import (
	"fmt"
	"strconv"
	"sync"
)

type SyncMap struct {
	sync.Mutex
	m map[string]int
}

var syncMap = SyncMap{m: make(map[string]int)}

func operation(k string, v int, wg *sync.WaitGroup) {
	defer wg.Done()
	syncMap.Lock()
	defer syncMap.Unlock()
	syncMap.m[k] = v
	fmt.Println(k, v)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go operation(strconv.Itoa(i), i, &wg)
	}
	wg.Wait()
}
