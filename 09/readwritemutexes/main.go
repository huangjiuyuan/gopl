package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go Deposit(100)
	go Deposit(200)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(Balance())

	go Withdraw(100)
	go Withdraw(100)
	go Withdraw(150)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(Balance())
}

var mu sync.RWMutex
var balance int

func Balance() int {
	mu.RLock() // readers lock
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

// This function requires that the lock be held.
func deposit(amount int) { balance += amount }
