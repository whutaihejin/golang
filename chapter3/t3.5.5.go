package main

import "fmt"
import "strconv"

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))
	s := fmt.Sprintf("%b", x)
	val, err := strconv.ParseInt("123", 10, 64)
	fmt.Printf("s=%s, val=%d, err=%s\n", s, val, "test")
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
}
