package main

import "fmt"
import "unicode/utf8"

func main() {
	s := "世界"
	s1 := "\xe4\xb8\x96\xe7\x95\x8c"
	s2 := "\u4e16\u754c"
	s3 := "\U00004e16\U0000754c"
	fmt.Printf("s=%s\n", s)
	fmt.Printf("s1=%s\n", s1)
	fmt.Printf("s2=%s\n", s2)
	fmt.Printf("s3=%s\n", s3)

	r := '\x41'
	fmt.Printf("r=%c\n", r)
	// r = '\xe4\xb8\x96'
	// r = '\Ue4b896'
	fmt.Printf("r=%c\n", r)

	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d %q %d\n", i, r, r)
	}

	r = 19990
	fmt.Printf("r=%c\n", r)

	n := 0
	for _, _ = range s {
		n++
	}
	fmt.Printf("n=%d\n", n)

	n = 0
	for range s {
		n++
	}
	fmt.Printf("n=%d\n", n)
	fmt.Printf("n=%d\n", utf8.RuneCountInString(s))

	name := "世界"
	fmt.Printf("% x \n", name)
	runeName := []rune(name)
	fmt.Printf("%x\n", runeName)
	fmt.Println(string(runeName))
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567))
}
