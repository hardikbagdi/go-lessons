package main

import (
	"encoding/json"
	"fmt"
)

type foo struct {
}
type bar interface {
}

func main() {
	var f foo
	var b bar
	buf, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf2, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
	fmt.Println(string(buf2))
	buf, _ = json.Marshal(nil)
	fmt.Println(string(buf))

}
