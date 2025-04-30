package main 
import "fmt"

const LoginToken string = "ghabbhhjd" // Public
func main() {
	var username string = "John Sena"
	var age int = 30

	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)
	fmt.Printf("Age is: %d\n", age)
	var isLoggedIn bool = true
	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	fmt.Print(isLoggedIn)


	// var anotherVariable int; //its not recomended to use this way 
	// anotherVariable = 10
	// fmt.Println(anotherVariable)

	// implicit type 
	var website = "Htttp://www.google.com"
	fmt.Print(website);
	// fmt.printf("\n")
	// no var 
	numberOfUsers := 563215564
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}

/*

| Feature               | `:=`                              | `=`                          |
|-----------------------|-----------------------------------|-------------------------------|
| Meaning               | Declare **and assign** a variable | Only **assign** a value       |
| Type inference        | Yes                               | No (requires variable already declared) |
| Scope                 | Only inside functions             | Can be used anywhere (where variable is declared) |
| Declaration required? | Yes (declares variable)           | No (just assigns)             |
| Example               | `x := 10`                         | `x = 20`                      |
*/