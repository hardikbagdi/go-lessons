package main

import "fmt"

func main() {

	fmt.Println("Enter a number: ")
	var input float32
	fmt.Scanf("%f", &input)
	

	//var input1 int
	//fmt.Scanf("%f", &input1) //surprisingly this compiles without a warning;why?

	output := input * 3
	fmt.Println(output)

}
