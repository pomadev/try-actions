package main

import "fmt"

// Add adds a and b, and returns the result
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Add(1, 1))
}
