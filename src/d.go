package main

import "fmt"

type Interface interface{}

var nilInterface Interface

func test() Interface {
	return nilInterface
}

func main() {
	fmt.Println("vim-go")

	x := test()
	fmt.Printf("\n%d\n", x)

}
