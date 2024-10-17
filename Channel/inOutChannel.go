package main

import (
	"fmt"
	"sync"
)

func inData(message chan<- string, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	message <- name

}

func outData(val <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	result := <-val

	fmt.Println(result)

}

func main() {
	var wg sync.WaitGroup
	data := make(chan string)

	wg.Add(2) // ini berarti menahan sebanyak 2 goroutine
	// Sehingga ketika done . bisa lanjut ke goroutine 1 lagi . dan wait masih tetap menunggu
	go inData(data, "Mikhel", &wg)

	go outData(data, &wg)

	wg.Wait()

}
