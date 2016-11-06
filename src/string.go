package main

import "fmt"

func main() {
	fmt.Println(len("Hello World!"))
	fmt.Println("Hello World"[1]) //int of e (char is represent as byte-ASCII)
	fmt.Println("Hello " + "World")
}
