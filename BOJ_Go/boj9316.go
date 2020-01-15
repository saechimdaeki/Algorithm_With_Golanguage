package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 1; i <= N; i++ {
		fmt.Printf("Hello World, Judge %d!\n", i)
	}
}
