package main

import "fmt"

type person struct {
	name string
	age  int
	food []string
}

func main() {
	food := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 25, food: food}
	fmt.Println(nico.name)
}
