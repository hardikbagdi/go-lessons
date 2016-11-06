package main

import "fmt"

func mul(x *int) {
	*x = *x * *x
}
func main() {
	x := new(int)
	*x = 2
	mul(x)
	fmt.Println(*x)
}
