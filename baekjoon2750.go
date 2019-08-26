package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	mySlice := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanln(&a)
		mySlice = append(mySlice, a)
	}
	sort.Ints(mySlice)
	for _, i := range mySlice {
		fmt.Println(i)
	}
}
