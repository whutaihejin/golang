package main

import "fmt"
import "strings"
import "bytes"

func main() {
	fmt.Printf("%s %s\n", "a/b/c.go", basename("a/b/c.go"))
	fmt.Printf("%s %s\n", "c.d.go", basename("c.d.go"))
	fmt.Printf("%s %s\n", "abc", basename("abc"))
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))

	s := "taihein"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s)
	fmt.Println(s2)

	fmt.Printf("Contains(taihejin, hejin) %t\n", strings.Contains("taihejin", "hejin"))
	fmt.Printf("Count(hello hello world hello, hello) %d\n", strings.Count("hello hello world hello", "hello"))
	fields := strings.Fields("hello world")
	fmt.Println(len(fields), fields)
	fmt.Println(strings.HasPrefix("taihejin", "tai"))
	fmt.Println(strings.Index("hello my name is", "my"))
	fmt.Println(strings.Join([]string{"hello", "taihejin", "tomhanks"}, "#"))

	fmt.Println(toString([]int{1, 2, 3}))
}

func basenameV1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func toString(nums []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range nums {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
