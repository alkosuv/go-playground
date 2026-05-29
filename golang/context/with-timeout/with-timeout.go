package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cansel := context.WithTimeout(context.Background(), time.Second*20)
	defer cansel()

	count := 0
	for {
		ch := ctx.Done()
		select {
		case <-ch:
			fmt.Println("timeout")
			return
		default:
		}

		count++
		fmt.Println(count)
		if count == 2 {
			time.Sleep(time.Second * 5)
		}
		if count == 5 {
			break
		}
	}
}
