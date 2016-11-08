package main

import (
	"os"
	"log"
	"io/ioutil"
	)

func main(){
	logger := log.New(os.Stdout,"HELLO",log.Lmicroseconds)
	logger.SetOutput(ioutil.Discard)
	logger.Println("Hello from log")
	log.Println("direct")
	

}
