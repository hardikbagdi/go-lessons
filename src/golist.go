package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	l.PushBack(3)
	l.PushFront("Ringo")
	x := l.Back()
	fmt.Println(x.Value)
	y := l.Front()
	fmt.Println(y.Value)
}
