package main

import "fmt"

func pinger(c chan<- string) {
	for 2 == 2 {
		c <- "ping"
	}
}

func printer(c <-chan string, name string) {
	for 2 == 2 {
		fmt.Println(name, <-c)
	}
}

func main() {

	var channel chan string = make(chan string)
	go pinger(channel)
	go printer(channel, "Printer1")
	go printer(channel, "Printer2")
	var input string
	fmt.Scanln(&input)
}
