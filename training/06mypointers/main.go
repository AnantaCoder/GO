package main

import "fmt"

/*pointers
A pointer is a variable that stores the memory address of another variable.

In Go:

& → address of operator

* → dereference (access the value stored at that address)
*/
func swap(a*int,b*int){
	*a,*b = *b,*a
}

func main(){
	
	x := 10
	ptr := &x
	fmt.Println("Value of ",ptr)
	fmt.Println("dereferencing",*ptr)
	*ptr = 50
	fmt.Println("Value of ",x)

	y := 48
	z := 43
	temp2 := &y
	temp3 := &z
	swap(temp2,temp3)
	fmt.Println("The value of x and y is : ",y,z)

}