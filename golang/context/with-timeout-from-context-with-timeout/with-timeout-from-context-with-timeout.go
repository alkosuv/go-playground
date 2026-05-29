package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cansel := context.WithTimeout(context.Background(), time.Second*3)
	defer cansel()

	SubFunc(ctx, 1)
	SubFunc(ctx, 4)
}

func SubFunc(ctx context.Context, second int) {
	ctx, cansel := context.WithTimeout(ctx, time.Second*time.Duration(second))
	defer cansel()

	time.Sleep(time.Second * time.Duration(second))

	select {
	case <-ctx.Done():
		log.Printf("context doen. timeout: %ds", second)
	default:
		log.Printf("func down. timeout: %ds", second)
	}
}
