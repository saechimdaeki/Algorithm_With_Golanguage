package main

import "fmt"

func main() {
	var a, b int
	for {
		fmt.Scanln(&a, &b)

		if a == 0 && b == 0 {
			return
		}
		fmt.Println(a + b)
	}
}
