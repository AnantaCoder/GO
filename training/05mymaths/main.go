package main

import (
	"fmt"
	"math/big"
	"time"
	"crypto/rand"
)

func main() {
	a := 10
	b := 3

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", int64(a)/int64(b)) // Type Casting to int64 for accurate division
	fmt.Println("Modulus (Remainder):", a%b)

	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2025, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))


	// generating random number using cryptography 
	max := big.NewInt(100) // range 0-99
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		fmt.Println("Error generating random number:", err)
	} else {
		fmt.Println("Random number (0-99):", n)
	}
}
