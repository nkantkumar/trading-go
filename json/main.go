package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	p := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Encode to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
