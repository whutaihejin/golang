package main

import "fmt"

func main() {
	p := new(int)
	fmt.Printf("p=%d\n", *p)
	*p = 2
	fmt.Printf("p=%d\n", *p)

	x := new(int)
	y := new(int)
	fmt.Printf("x == y %t\n", x == y)

	px := newInt()
	py := newInteger()
	fmt.Printf("px=%x, *px=%d, py=%x, *py=%d\n", px, py, *px, *py)
}

func newInt() *int {
	return new(int)
}

func newInteger() *int {
	var value int
	return &value
}
