package main

import "fmt"

var global *int

func f() {
	var x int = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}

func main() {
	fmt.Printf("global=%v\n", global)
	f()
	g()
	fmt.Printf("global=%v\n", global)
}
