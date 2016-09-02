package main

import "fmt"
import "time"

const (
	a = 1
	b
	c = 2
	d
)

func main() {
	const pi = 3.1415926
	fmt.Printf("%T %v\n", pi, pi)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %v\n", noDelay, noDelay)
	fmt.Printf("%T %v\n", timeout, timeout)
	fmt.Printf("%T %v\n", time.Minute, time.Minute)

	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)
}
