package main

import "fmt"

// go lang tutorial 연습문제 풀이
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		defer func() {
			a = a + b
			a, b = b, a
		}()
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
