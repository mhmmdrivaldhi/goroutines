package main

import (
	"fmt"
	"sync"
	"time"
)

// var wg sync.WaitGroup - // jika global variable tidak perlu menambahkan parameter di func lain.

func main() {
	// fmt.Println("One")
	// fmt.Println("Two")
	// time.Sleep(time.Second * 3)
	// fmt.Println("Three")
	// fmt.Println("Done")
	var wg sync.WaitGroup // jika local variable maka harus menambahkan parameter baru di func lain dan juga menggunakan pointer
	wg.Add(2)

	fmt.Println("Start")
	go repeat("Ping", 500, &wg)
	go repeat("Signal Network", 500, &wg)
	wg.Wait()
	fmt.Println("Done")


}

func repeat(word string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		fmt.Println(i, word)
		time.Sleep(time.Millisecond * delay)
	}
} 