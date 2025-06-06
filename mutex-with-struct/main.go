package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	account := BankAccount{Balance: 0}

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		account.topUp(10)
	}
	wg.Wait()
	fmt.Println(account.Balance)
}

type BankAccount struct {
	Balance int
	mutex sync.Mutex
}

func (account *BankAccount) topUp(amount int) {
	defer account.mutex.Unlock() 
	defer wg.Done()
	account.mutex.Lock()
	account.Balance += amount
}