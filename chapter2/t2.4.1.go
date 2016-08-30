package main

import "fmt"

func main() {
	x, y := 1, 2
	fmt.Printf("x = %d, y = %d\n", x, y)
	x, y = y, x
	fmt.Printf("x = %d, y = %d\n", x, y)
	x, y = 11, 12
	fmt.Printf("x = %d, y = %d\n", x, y)
}
