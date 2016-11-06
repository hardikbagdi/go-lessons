package main

import "fmt"


func test() (id int) {
	id = 42
	return
}


func foo() ( int ,int, int , bool) {

	return 0,-1,-2,true
}


func bar () func() (int ,int ,int, bool) {

retfun := func() ( int ,int, int , bool) {

	return 42,-11,-22,true
}
return retfun
}

func main(){
	x:= 0
	//a closure
	increment := func() int {
		x++	
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(test())
	fmt.Println(foo())
	foohere := bar()
	fmt.Println(foohere())
}
