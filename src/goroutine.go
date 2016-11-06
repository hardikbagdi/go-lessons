package main

import "fmt"
import "time"
import "math/rand"

func foo(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Goroutine:", i)
		go foo(i)
	}
	var input string
	fmt.Scanln(&input)
}
