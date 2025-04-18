package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza shop")
	fmt.Println("please rate btn 1 and 10")
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString(('\n')) //read up to  the new line , _ is for the error handling
	input = strings.TrimSpace(input) //removes both \r and \n (and any leading/trailing spaces)
	fmt.Println("Thanks for the rating of ",input)
	numRating,err := strconv.ParseFloat(input,64)

	if err!=nil{
		fmt.Println("error",err)
	}

	if numRating<1 || numRating>10{
		fmt.Println("Please enter a number between 1 and 10")
	} else if numRating >= 1 && numRating <= 3 {
		fmt.Println("We are sorry to hear that you had a bad experience")
	} else if numRating >= 4 && numRating <= 6 {
		fmt.Println("We are glad to hear that you had an average experience")
	} else if numRating >= 7 && numRating <= 8 {
		fmt.Println("We are happy to hear that you had a good experience")
	} else if numRating >= 9 && numRating <= 10 {
		fmt.Println("We are thrilled to hear that you had an amazing experience")
	}

}