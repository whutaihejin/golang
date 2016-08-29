package main

import "fmt"

const boilingF = 212.0

func add(x, y int32) int32 {
	return x + y
}

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g℉ or %g℃\n", f, c)
}
