package main

import (
	"fmt"
)

func main() {
	fmt.Print("array: ")

	var fruitlist []string
	for i := 0; i < 5; i++ {
		fruitlist = append(fruitlist, fmt.Sprintf("%d", i))
	}
	
	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist))

    var vegList = [5]string{"potato","tomato","bonana","kokoa"}
	fmt.Print(vegList)
	fmt.Print(len(vegList))
}
