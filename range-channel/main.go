package main

import "fmt"

func main() {
	channel := make(chan int)

	go func(){
		defer close(channel)
		channel <- 7
		channel <- 5
		channel <- 4
	}()

	avarage(channel)
}

func avarage(channel chan int) {
	var sum = 0

	for value := range channel {
		sum = sum + value
		fmt.Println("Avarage : ", sum)
	}
}