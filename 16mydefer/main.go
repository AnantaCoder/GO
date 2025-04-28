package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()

}

// world, One, Two
// 0, 1, 2, 3, 4
// hello, 43210, two, One, world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

/**
🔥 Important Concept about defer in Go:
defer statements delay the execution of a function or statement until the surrounding function (main or myDefer) returns.

Deferred calls are pushed onto a stack — meaning the last defer is executed first (LIFO: Last-In, First-Out).
*/