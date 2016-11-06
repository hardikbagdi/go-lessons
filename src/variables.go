package main

import "fmt"

func main() {
	//var <name> <data-type>
	var x string = "Hello World" //string is the datatype
	fmt.Println(x)

	var y int
	y = 42
	fmt.Println(y)

	//shorthand declaration
	//notice : before =  for better code readability( indicate declaration)
	//type inference done by go compiler
	z := "Bye World"
	fmt.Println(z)
	
	var (
		n = 42
		m = 3.14
	)
	fmt.Println(n)
	fmt.Println(m)

	var (
		n1 string
		m1 bool
	)
	fmt.Println(n1)
	fmt.Println(m1)

	var (
		n2 string = "Ola"
		m2 float32 = 1.618
	)
	fmt.Println(n2)
	fmt.Println(m2)
}
