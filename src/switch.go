package main

import "fmt"
import "math/rand"

func main() {
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		switch x % 2 {

		case 0:
			fmt.Println(x, " is even")
		case 1:
			fmt.Println(x, " is odd")
		default:
			fmt.Println("Great Scott")
		}
	}
}
