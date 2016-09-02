package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

	fmt.Printf("%b %b %b %b %b\n", FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

	fmt.Printf("%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d", KiB, MiB, GiB, TiB, PiB, EiB, ZiB, YiB)
}
