package main

import "fmt"

func main() {
	fmt.Printf("Go prime ")

	var i  int
	fmt.Scanf("%d", &i)
	for i := 0; i < 5; i++ {		
		fmt.Println(i)
	}

	prime(i)
}

func prime(num int) {
	if num <= 0 {
		fmt.Println("Not Prime")
		return
	}
	isPrime := true
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime && num > 1 {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not prime")
	}
}