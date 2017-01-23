package main

import (
	"fmt"
	"os"

	"github.com/google/go-querystring/query"
)

type Options struct {
	Query   string `url:"q"`
	ShowAll bool   `url:"all"`
	Page    int    `url:"page"`
}

type arr struct {
	Num []int `url:"nums"`
}

type st struct {
	Num string `url:"nums"`
}

func main() {
	opt := Options{"foo", true, 2}
	v, _ := query.Values(opt)
	fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
	getQS(opt)
	array := []int{1, 2, 3, 4}
	var a arr
	a.Num = array
	v, err := query.Values(a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(v.Encode())

	s := "1,2,3,4"
	var s1 st
	s1.Num = s
	v, err = query.Values(s1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(v.Encode())
}

func getQS(foo1 interface{}) {
	fmt.Println("getQs")
	v, _ := query.Values(foo1)
	fmt.Println(v.Encode())
}
