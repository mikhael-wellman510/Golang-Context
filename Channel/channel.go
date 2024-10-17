package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		message <- "Mikhael"

	}()

	result1 := <-message
	fmt.Println(result1)

}
