package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subtract(x int ,y int)int{
	return x - y
}
func multiply(x int,y int )int{
	return x + y
}
func proadder(values ...int)(int,string){
	sum := 0
for _,value := range values{
	sum += value 
}
return sum,"The sum of the values is : "
}



func main() {
	result := add(3, 5)
	fmt.Println("The sum is:", result)
	result1 := subtract(10,4)
	fmt.Println("The difference is :", result1)
    result2 := multiply(54,54)
	fmt.Println("The product is :", result2)
	numbers,message := proadder(1,2,3,4,5,54,3,32,2,2,4)
	fmt.Println(message,numbers)
}