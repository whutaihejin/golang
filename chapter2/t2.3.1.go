package main

import "fmt"

func main() {
	i, j := 2, 9
	fmt.Printf("i=%d, j=%d\n", i, j)
	i, j = j, i
	fmt.Printf("i=%d, j=%d\n", i, j)
	x, y := 1, 2
	fmt.Printf("x=%d, y=%d\n", x, y)
	z, y := 5, 6
	fmt.Printf("z=%d, y=%d\n", z, y)
	m, n := 8, 9
	fmt.Printf("m=%d, n=%d\n", m, n)
	// m, n := 11, 22 // no new variables on left side of :=
	fmt.Printf("m=%d, n=%d\n", m, n)
}
