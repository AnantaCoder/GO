package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func greetUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s! Welcome!\n", name)
}

func main() {
	greetUser()
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		// Loop logic here (currently empty)
		if i == 999999{
			fmt.Printf("Loop Completed ")
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}