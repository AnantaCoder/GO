package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Files is golang ")
	content := "This is a file which will be inside the golang project "
	os.Create("./17files/myfile.txt")
	file, err := os.OpenFile("./myfile.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error while opening the file", err)
	} else {
		fmt.Println("File opened successfully")
	}
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error while writing to the file", err)
	} else {
		fmt.Println("File written successfully")
	}
	// file.Close()
	// file.Close() is not needed as it is done automatically by the defer statement
	defer file.Close()
	fmt.Println("File closed successfully")

}