package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan int, 3)

	wg.Add(3)

	go send(channel, 7)
	go send(channel, 6)
	go send(channel, 5)


	// -- on one function

	// channel <- 1
	// channel <- 2
	// channel <- 3
	// fmt.Println("Lenght Channel : ", len(channel))
	// fmt.Println(<- channel)
	// fmt.Println(<- channel)
	// fmt.Println(<- channel)
	// fmt.Println("Capacity Channel : ", cap(channel))
	wg.Wait()
	sum(channel)
	fmt.Println("Done")
}

func send(channel chan int, number int) {
	defer wg.Done()
	channel <- number
}

func print(channel chan int){
	defer wg.Done()
	fmt.Println(<- channel)
}

func sum(channel chan int) {

	result := 0

	for len(channel) > 0 {
		result += <-channel
	}
	fmt.Println(result)
}