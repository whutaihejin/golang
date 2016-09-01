package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Printf("len(s)=%d\n", len(s))
	fmt.Println(s[0], s[7])
	// panic
	// c := s[len(s)]
	// c++
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("goodbye" + s[5:])

	v := "left root"
	t := v
	v += ", right foot"
	fmt.Println(v)
	fmt.Println(t)

	// cannot assign to s[0]
	// s[0] = 'L'
}
