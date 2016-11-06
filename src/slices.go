package main

import (
	"fmt"
)

func main() {
	x := make([]int, 4, 10)
	x[0] = 11
	x[1] = 22
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)
	//x[5] = 10; //panic: index out of range
	a := []float64{12, 3} //this also is a slize
	fmt.Println(cap(a))
	fmt.Println(len(a))
	b := [3]int{1, 2, 3}
	fmt.Println(cap(b)) //cap(array) will give the length
	fmt.Println(len(b))

	//slicing a slice
	y := x[2:7]
	y[0] = -12 //this edits x[2]
	//slices backed by the same array
	fmt.Println(x)
	fmt.Println(y)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice1[0] = 11
	fmt.Println("Modified slice1")
	fmt.Println(slice1)
	fmt.Println(slice2)
}
