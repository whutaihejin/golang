package main

import "fmt"

func main() {
	s := "Hello, 世界"
	fmt.Println(s)
	const GoUsage = `Go is a tool for managing Go source code.
	
	Usage:
		go command [arguments]
	...`
	fmt.Println(GoUsage)
}
