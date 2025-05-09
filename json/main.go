package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"e-mail"`
}

func main() {
	p1 := Person{Name: "Ale", Age: 25, Email: "alexandre@mail.com"}
	fmt.Println("Person 1 struct:", p1)

	p1Json, _ := json.Marshal(p1)
	fmt.Println("Person 1 JSON:", string(p1Json))

	p1Map := make(map[string]any)
	json.Unmarshal(p1Json, &p1Map)
	fmt.Println("--- Person 1 map ---")
	for key, value := range p1Map {
		fmt.Println(key, ":", value)
	}
	fmt.Println()
	
	var p1Person Person
	json.Unmarshal(p1Json, &p1Person)
	fmt.Println("Person 1 struct again:", p1Person)
}
