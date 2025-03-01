package main

import "fmt"

func increment(n *int) {
	*n = *n + 1 // Dereferencing the pointer to modify original value
}

func main() {
	x := 5
	increment(&x)  // Pass address of x
	fmt.Println(x) // 6 (original x modified)
}
