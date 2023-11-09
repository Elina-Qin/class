package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func subtract(x int, y int) int {
	return x - y
}
func multiply(x int, y int) int {
	return x * y
}
func main() {
	x := 7
	y := 11
	fmt.Printf("Add:%d\n", add(x, y))
	fmt.Printf("Subtract:%d\n", subtract(x, y))
	fmt.Printf("Multiply:%d\n", multiply(x, y))
}
