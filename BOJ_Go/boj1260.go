package main

import "fmt"

var N, M, V int
var grid [1001][1001]int
var visited [1001]bool
var isFirst bool = true

func dfs(start int) {
	if visited[start] {
		return
	}
	visited[start] = true
	if isFirst == true {
		isFirst = false
	} else {
		fmt.Print(" ")
	}
	fmt.Print(start)
	for i := 1; i <= N; i++ {
		if grid[start][i] == 1 && false == visited[i] {
			dfs(i)
		}
	}
}

func bfs(start int) {
	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if visited[cur] == true {
			continue
		}
		visited[cur] = true
		if isFirst {
			isFirst = false
		} else {
			fmt.Print(" ")
		}
		fmt.Print(cur)
		for i := 1; i <= N; i++ {
			if grid[cur][i] == 1 && visited[i] == false {
				queue = append(queue, i)
			}
		}
	}

}

func main() {

	fmt.Scanln(&N, &M, &V)
	for i := 0; i < M; i++ {
		var start, end int
		fmt.Scanln(&start, &end)
		grid[start][end] = 1
		grid[end][start] = 1
	}
	for i := 1; i <= N; i++ {
		visited[i] = false
	}
	isFirst = true
	dfs(V)

	fmt.Println()

	for i := 1; i <= N; i++ {
		visited[i] = false
	}
	isFirst = true
	bfs(V)

}
