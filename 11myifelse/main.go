package main

import "fmt"

func main() {
	fmt.Println("If else go")
	var m int
	fmt.Print("Enter a number: ")
	fmt.Scan(&m)
	if m%2 == 0 {
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}
}