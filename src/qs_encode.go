package main

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type Options struct {
	Query   string `url:"q"`
	ShowAll bool   `url:"all"`
	Page    int    `url:"page"`
}

func main() {
	opt := Options{"foo", true, 2}
	v, _ := query.Values(opt)
	fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
	getQS(v)
}

func getQS(foo1 interface{}) {
	fmt.Println("getQs")
	v := foo1.(url.Values)
	fmt.Println(v.Encode())
}
