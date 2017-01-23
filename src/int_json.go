package main

import "fmt"
import "encoding/json"

type wrapper struct {
	I *int `json:"int"`
	J int  `json:"val"`
}

func main() {
	var w wrapper
	if w.I == nil {
		fmt.Println("is nill")
	}
	i := 42
	w.I = &i
	var w2 wrapper
	jsonString := []byte(`{"val": 2}`)
	st, _ := json.Marshal(w)
	fmt.Println("Marshalled : ", string(st))

	_ = json.Unmarshal(jsonString, &w2)
	fmt.Println(w2)
	if w2.I == nil {
		fmt.Println("Is nill")
	}
}
