package main

import "fmt"

type Book struct {
	Uid int64
	Lng float64
	Lat float64
}

func (book *Book) shift() {
	if book == nil {
		return
	}
	book.Uid += 1
	book.Lng += 1
	book.Lat += 1
}

func (book *Book) unshift() {
	if book == nil {
		return
	}
	book.Uid -= 1
	book.Lng -= 1
	book.Lat -= 1
}

func (book Book) String() string {
	return fmt.Sprintf("addr=%p, uid=%d, lng=%g, lat=%g", &book, book.Uid, book.Lng, book.Lat)
}

func main() {
	book := Book{1, 22, 33}
	var pb *Book = &book
	pb.shift()
	fmt.Println(book)
	pb.unshift()
	fmt.Println(book)
	//
	pb = nil
	pb.shift()
	fmt.Println(book)
	pb.unshift()
	fmt.Println(book)
	//
	book.shift()
	fmt.Println(book)
	book.unshift()
	fmt.Println(book)
	//
	var bookNil Book
	fmt.Println(bookNil)
}
