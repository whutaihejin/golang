package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Printf("len(s)=%d\n", len(s))
	fmt.Println(s[0], s[7])
	// panic
	c := s[len(s)]
	c++
}
