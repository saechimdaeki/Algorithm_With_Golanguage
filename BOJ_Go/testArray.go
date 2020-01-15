package main

import "fmt"

func main() {
	arr := []int{0, 1}

	arr = append(arr, 6, 7, 8, 9)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}
