package main

import "fmt"

func main() {

	nico := map[string]string{"name": "nico", "age": "12"}
	//fmt.Println(nico)
	for _, value := range nico {
		fmt.Println(value)
	}
	/* output:
	nico
	       12
	*/
}
