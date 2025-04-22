package main

import (
	"encoding/json"
	"fmt"
)

// structs - grouping of related data , collection of fields
type Person struct {
	name    string
	age     int
	id      int
	isAlive bool
	Address // embedded structs composition
}
type Address struct {
	City    string
	State   string
	Country string
}

// updating by reference
func updateName(p *Person, newName string) {
	p.name = newName
}

type newAddress struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	newAddress     // embedded composition
}

func main() {
	fmt.Print("Struct in go ")
	a1 := Address{"New York", "NY", "USA"}
	a2 := Address{"Los Angels", "LA", "USA"}
	a3 := Address{"Chicago", "CK", "USA"}
	var p1 = Person{"John", 30, 1, true, a1}
	p2 := Person{"Kaka", 30, 1, true, a2}
	p3 := Person{"Raheem", 30, 1, false, a3}

	updateName(&p1, "Fko wakei")
	fmt.Println("Name of p2:", p2.name)
	fmt.Println("Is p3 alive?", p3.isAlive)
	fmt.Println(p1)
	fmt.Printf("p2 details along with the parameters are: %+v\n", p2)
	fmt.Println("P1 ADDRESS", p1.Address)

	user := User{
		Name:  "Anirban",
		Email: "anirban@example.com",
		Age:   21,
		newAddress: newAddress{
			City:    "Kolkata",
			State:   "West Bengal",
			Country: "India",
		},
	}
	// Convert struct to JSON
	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(jsonData))
}
