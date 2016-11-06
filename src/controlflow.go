package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else { //you can't skip the braces if there is one statement like you can do in C/Java
			fmt.Println(i, " is odd")
		}
	}

}
