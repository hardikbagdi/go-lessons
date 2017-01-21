package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Twrapper struct {
	Time time.Time `json:"time"`
}

func main() {
	t := " [{\"time\": \"2013-06-29T09:42:48-04:00\"},{\"time\": \"2013-06-29T09:57:48-04:00\"}]"
	tval := make([]Twrapper, 0)
	err := json.Unmarshal([]byte(t), &tval)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(tval)
}
