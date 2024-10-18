package main

import (
	"context"
	"fmt"
	"time"
)

func fetch(ctx context.Context) {

	for i := 0; i < 5; i++ {
		select {
		// Jika waktu terlalu lama . maka akan langsung stop
		case <-ctx.Done():
			fmt.Println("Failed , Waktu terlalu lama")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Data : ", i)
		}
	}

	fmt.Println("Succes Get data")
}

func main() {

	Parent := context.Background()

	ctx, cancel := context.WithTimeout(Parent, 3*time.Second)
	defer cancel()

	fetch(ctx)

	time.Sleep(3 * time.Second)

}
