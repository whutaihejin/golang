package main

import "fmt"
import "strconv"

func strconvTest(str string) {
	defer fmt.Printf("==============\n")
	fmt.Printf("str %s\n", str)
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("err %s\n", err)
		return
	}
	fmt.Printf("val %d type %T err %s\n", val, val, err)
}

func main() {
	strconvTest("")
	strconvTest("12")
	strconvTest("+12")
	strconvTest("-12")
	strconvTest("0")
	strconvTest("0000")
	strconvTest("11abc")
	strconvTest("abc")
}
