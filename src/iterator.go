package main

import "fmt"

func iterator(array []int) chan int {
	ch := make(chan int, 2)
	go func(arr []int) {

		for val := range arr {
			ch <- val
		}
		close(ch)
	}(array)
	return ch
}
func main() {
	fmt.Println("vim-go")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for val := range iterator(arr) {

		fmt.Println(val)
	}

}
