package main

import "fmt"

type Person struct{
	name string
}
type Student struct{
	Person
	id int
	major string
}
func (s *Student) getName() string {
	return s.name
}
func main(){
	s := new (Student)
	s.id = 42
	s.name = "Cosmo Kramer"
	s.major = "Proctology"
	fmt.Println(*s) 
	fmt.Println(s) 
	fmt.Println(s.getName())
}
