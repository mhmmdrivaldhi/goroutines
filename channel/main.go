package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan string)
	defer close(channel)

	wg.Add(2)
	// go gateway(channel, "Hi, Rivaldhi!")

	// fmt.Println(<- channel)

	go sendMessage(channel, "What's up Bro ?")
	go getMessage(channel)

	wg.Wait()
	fmt.Println("Done")
}

func gateway(channel chan string, words string) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	channel <- words
} 

func sendMessage(channel chan <- string,  message string) {
	defer wg.Done()
	channel <- message
}

func getMessage(channel <- chan string) {
	defer wg.Done()
	fmt.Println(<- channel)
}