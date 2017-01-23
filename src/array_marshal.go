package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type foo1 struct {
	A int `json:"a"`
}

func main() {
	//	st1 := `{ "a" : 10}`
	//st2 := `[{"a" : 10}]`
	st3 := `[{"a" : 10},{"a" : 20}]`

	arr := make([]foo1, 0)
	err := json.Unmarshal([]byte(st3), &arr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arr)
}
