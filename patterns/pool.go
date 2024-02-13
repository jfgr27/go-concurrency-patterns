package main

import (
	"fmt"
	"sync"
	"time"
)

type PoolWork struct {
	w int
	t int
}

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Millisecond)
		results <- task * 2
	}
}

// Runs w workers concurrently.
// Each will pick work on all tasks, doubling value.
// t tasks are added to tasks channels.
// A worker id will pick up, increment and write to results

// Note: if t >> w, workers will process slowly
// Note if w >> t, workers won't all be used
func (c *PoolWork) work() {
	s := fmt.Sprintf("Pool with %d tasks, %d workers", c.t, c.w)
	defer timeit(s)()

	tasks := make(chan int, c.t)
	results := make(chan int, c.w)

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= c.w; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, tasks, results)
		}(i)
	}

	// Enqueue jobs
	for i := 1; i <= c.t; i++ {
		tasks <- i
	}
	close(tasks)

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()


	// Collect results
	var res int
	for s := range results {
		res += s
	}

	fmt.Println("Output: ", res)
}
