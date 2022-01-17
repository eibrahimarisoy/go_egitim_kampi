package main

import "fmt"

func main() {
	//channel initialization
	unBufferedChan := make(chan int)

	// channel declaration
	var unBufferedChan2 chan int

	fmt.Println(unBufferedChan)
	fmt.Println(unBufferedChan2) //nil

	go func(unBufChan <-chan int) {
		// blocks until data
		value := <-unBufChan
		fmt.Println(value)

	}(unBufferedChan)

	unBufferedChan <- 1
}
