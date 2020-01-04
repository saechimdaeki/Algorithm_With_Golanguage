package main

import "fmt"

func solution(x int) int {
	sum := x
	for {
		if x == 0 {
			break
		}
		sum += x % 10
		x = x / 10
	}
	return sum
}

var visited [10002]bool

func main() {
	for i := 1; i < 10001; i++ {
		var abc = solution(i)
		if abc <= 10000 {
			visited[abc] = true
		}
	}
	for i := 1; i < 10000; i++ {
		if visited[i] == false {
			fmt.Println(i)
		}
	}
}
