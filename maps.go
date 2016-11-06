package main

import "fmt"

func main(){

	//need to instantiate before using;else panic
	//var x map[string]int
	//x["Hello"] = 42
	//fmt.Println(x)
	x := make(map[string]int)
	x["hello"] = 42
	value, ok := x["hello"]
	fmt.Println(value,ok)
	fmt.Println(x)
	test := map[string]int{
		"a" : 121,
		"b" : 23, //notice the ,
	}
	test["c"] = 312
	fmt.Println(test)
}
