package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	req1, err := http.NewRequest("GET", "http://example.com", nil)
	req1.Write(os.Stdout)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := []byte("{}")
	req2, err := http.NewRequest("GET", "http://example.com", bytes.NewReader(buf))
	if err != nil {
		fmt.Println(err)
		return
	}
	req2.Write(os.Stdout)
	buf = []byte("null")
	req2, err = http.NewRequest("GET", "http://example.com", bytes.NewReader(buf))
	if err != nil {
		fmt.Println(err)
		return
	}
	req2.Write(os.Stdout)
}
