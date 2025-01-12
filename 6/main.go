package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	go func(ctx context.Context) {
		i := 0
		for {
			time.Sleep(time.Second)
			fmt.Println(i)
			i++
		}
	}(ctx)

	sec := rand.Int()%5 + 3
	fmt.Println(sec)

	time.Sleep(time.Duration(sec) * time.Second)
	cancel()
}
