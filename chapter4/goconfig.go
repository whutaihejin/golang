package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
)

const (
	section_test = "section_test"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("test.ini")
	if err != nil {
		log.Fatalf("can't load config file of test.ini: %s", err)
	}
	value1, err := cfg.GetValue(section_test, "key1")
	if err != nil {
		log.Fatalf("can't extract key1's value: %s", err)
	}
	fmt.Printf("%s = %s\n", "key1", value1)
	value2, err := cfg.GetValue(section_test, "key2")
	if err != nil {
		log.Fatalf("can't extract key2's value: %s", err)
	}
	fmt.Printf("%s = %s\n", "key2", value2)
}
