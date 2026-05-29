package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, jobs <-chan int, results chan<- int, numWorker int) {
	fmt.Printf("Worker %d started\n", numWorker)
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: Channel closed\n", numWorker)
				return
			}
			fmt.Printf("Worker %d: Processing job %d\n", numWorker, job)
			results <- job * 2
		}
	}
	fmt.Printf("Worker %d finished\n", numWorker)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int)

	for i := 0; i < numWorkers; i++ {
		go worker(ctx, jobs, results, i)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for j := 0; j < numJobs; j++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}
