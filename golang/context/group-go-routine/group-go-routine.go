package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"log"
	"time"
)

func main() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		wg          sync.WaitGroup

		serviceNames = []string{"yandex", "uber", "maxim", "go taxi"}

		resultCh = make(chan string)
		winner   string
	)

	wg.Add(len(serviceNames))
	for _, s := range serviceNames {
		go func(serviceName string) {
			defer wg.Done()
			requestRide(ctx, serviceName, resultCh)
		}(s)
	}

	go func() {
		winner = <-resultCh
		// clouse channel
		close(resultCh)
		// Canceling context
		cancel()
	}()

	wg.Wait()

	fmt.Printf("found car in %q\n", winner)
}

func requestRide(ctx context.Context, serviceName string, resultCh chan<- string) {
	time.Sleep(time.Second * 3)

	for {
		select {
		case <-ctx.Done():
			log.Printf("stopped the search in %q (%v)\n", serviceName, ctx.Err())
			return

		default:
			if rand.Float32() > 0.75 {
				resultCh <- serviceName
				return
			}

			continue
		}
	}
}
