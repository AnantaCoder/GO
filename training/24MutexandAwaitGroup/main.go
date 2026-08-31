package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("race conditions")

	var score = []int{0}
	var mu sync.Mutex // for safe concurrent access

	wg := &sync.WaitGroup{}
	wg.Add(3) // Add 3 goroutines

	// Goroutine 1
	go func() {
		fmt.Println("one r")
		mu.Lock() //entry section 
		score = append(score, 1) //critical section 
		mu.Unlock() // exit section 
		wg.Done()
	}()

	// Goroutine 2
	go func() {
		fmt.Println("two r")
		mu.Lock()
		score = append(score, 2)
		mu.Unlock()
		wg.Done()
	}()

	// Goroutine 3
	go func() {
		fmt.Println("three r")
		mu.Lock()
		score = append(score, 3)
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final score is", score)
}
//A Mutex (Mutual Exclusion) is used to safely share data across multiple goroutines. It ensures that only one goroutine can access a critical section of code at a time.