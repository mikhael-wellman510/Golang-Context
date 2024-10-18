package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			// Jika sudah melewati waktu . dia akan return
			case <-ctx.Done():
				return
			default:
				// Memasukan counter ke chanel
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func main() {
	// memastikan goroutine tidak akan berjalan lagi jika melebihi batas waktu yang ditentukan
	fmt.Println("Total goroutine : ", runtime.NumGoroutine())

	// instance
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)

	// ini akan mengirim sinyal cancel ke context
	defer cancel()

	destination := CreateCounter(ctx)

	for val := range destination {
		fmt.Println(val)
	}

	// Memastikan Goroutine jumlah nya kembali ke semula
	fmt.Println("Total goroutine : ", runtime.NumGoroutine())

}
