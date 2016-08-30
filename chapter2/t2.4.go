package main

import "fmt"

func main() {
	var x int
	x = 1
	var p *int = new(int)
	*p = 2
	fmt.Printf("x = %d, p = %v, *p = %d\n", x, p, *p)
	x--
	x--
	fmt.Printf("x = %d, p = %v, *p = %d\n", x, p, *p)
	x++
	fmt.Printf("x = %d, p = %v, *p = %d\n", x, p, *p)
}
