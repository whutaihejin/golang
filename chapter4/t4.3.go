package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["taihejin"] = 11
	ages["tomhanks"] = 22
	for name, age := range ages {
		fmt.Printf("%s %d\n", name, age)
	}

	whiteList := make(map[string]bool)
	whiteList["280"] = true
	whiteList["281"] = true
	if whiteList["280"] {
		fmt.Printf("280 in the map\n")
	}
	if whiteList["260"] {
		fmt.Printf("260 not in the map")
	}
}
