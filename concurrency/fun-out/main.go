package main

import (
	"fmt"
	"time"
)

func main() {
	const numOperations = 3
	ch := make(chan struct{})
	defer close(ch)

	for i := 0; i < numOperations; i++ {
		go func(index int) {
			fmt.Printf("worker %d: start\n", index)
			time.Sleep(time.Second * 2)
			fmt.Printf("worker %d: done\n", index)
			ch <- struct{}{}
		}(i)
	}

	for count := 0; count < numOperations; count++ {
		<-ch
	}

	fmt.Println("main: finished")
}
