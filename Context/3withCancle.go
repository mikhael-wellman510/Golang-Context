package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func createCounter(ctx context.Context, wg *sync.WaitGroup) chan int {
	defer wg.Done()
	destination := make(chan int)

	go func() {

		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func main() {
	// Start goroutine cmn 1
	// ketika ada go func , maka go routine nambah 1
	var wg sync.WaitGroup
	fmt.Println("total goroutine ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	wg.Add(1)
	destination := createCounter(ctx, &wg)

	fmt.Println("Total go : ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter : ", n)

		if n == 10 {
			break
		}
	}

	cancel()
	wg.Wait()
	fmt.Println("Total goroutine : ", runtime.NumGoroutine())

}
