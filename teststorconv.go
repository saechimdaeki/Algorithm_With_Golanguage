package main

import (
	"fmt"
	"os"
	"strconv"
)

var s string

func main() {
	fmt.Scan(&s)
	// string to int
	test := 123
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(s, i+test)

}
