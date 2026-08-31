package main

import "fmt"

// maps - key and value pairs
func main()  {
	fmt.Println("Maps in go")
	language := make(map[string]string) //empty map 
	language["go"]="golang"
	language["py"] = "python"
	language["js"] = "javascript"

	fmt.Println("Languages are:  ",language)
	fmt.Println("js stands for:  ",language["js"])
	delete(language,"js")
	fmt.Println("Languages are:  ",language)

	//loops are 
	for key,value :=range language{
		fmt.Println("key:",key,"value:",value)
	}
	// var m map[int]string           // nil map
    var m = make(map[int]string)      // initialize the map
	for i:=0;i<=10;i++{
		m[i]= fmt.Sprintf("value %d",i);
	}
	for k,v :=range m{
		fmt.Println("key:",k,"value:",v)
	}


}