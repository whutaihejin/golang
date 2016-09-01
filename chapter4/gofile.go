package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main() {
	fileRead("nothing")
	fileRead("./test.ini")
}

func fileRead(filePath string) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("can't open %s: %s\n", filePath, err)
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	record, err := reader.ReadString('\n')
	for ; err == nil; record, err = reader.ReadString('\n') {
		record = strings.TrimSpace(record)
		fmt.Printf("record: %s\n", record)
	}
	return err
}
