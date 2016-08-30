package main

import "fmt"
import "flag"
import "strings"

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Printf(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	x := 1
	p := &x
	fmt.Printf("x=%d\n", *p)
	*p = 2
	fmt.Printf("x=%d\n", *p)

	var a, b int
	fmt.Printf("%t %t %t\n", &a == &a, &a == &b, &b == nil)

	vv := fun()
	fmt.Printf("%t\n", vv == fun())

	v := 1
	incr(&v)
	fmt.Printf("v=%d\n", incr(&v))
}

func fun() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
