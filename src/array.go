package main

import "fmt"
import "math/rand"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	total := 0
	for i := 0; i < len(x); i++ {
		x[i] = rand.Intn(100)
	}
	fmt.Println(x)
	for _, value := range x {
		total += value
	}
	fmt.Println("Average is ", float64(total)/float64(len(x)))
	//another
}
