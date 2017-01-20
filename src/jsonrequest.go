package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	buf := []byte("foo")
	fmt.Println(string(buf))
	req, err := http.NewRequest("GET", "http://example.com", bytes.NewReader(buf))
	if err != nil {
		fmt.Println(err)
		return
	}
	//req.Request.Header.Add("Content-Type", "application/json")
	req.Write(os.Stdout)
}
