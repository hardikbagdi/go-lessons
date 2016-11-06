package main

import "fmt"

func main() {

	fmt.Println("Hi American, enter a temperature in Fahrenheit: ")
	var input float32
	fmt.Scanf("%f", &input)

	//output = (input -32) * 5/9 //gives a compile error
	output := (input -32) * 5/9
	fmt.Println(output)

}
