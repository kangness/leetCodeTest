package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	ID   int
}

func ShowPerson(p Person) {
	p.Name = "Jia"
	fmt.Printf("in show person %+v\r\n", p)
	fmt.Printf("show person p address %p\r\n", &p)
}
func main() {
	p := Person{Name: "kangness", Age: 10, ID: 10}
	fmt.Printf("%+v --- --- \r\n", p)
	fmt.Printf("show person pp address %p\r\n", &p)
	ShowPerson(p)
	fmt.Printf("show person pp address %p\r\n", &p)
}
