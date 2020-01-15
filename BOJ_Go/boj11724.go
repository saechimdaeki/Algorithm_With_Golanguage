package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, M, V int
var grid [1001][1001]int
var visited [1001]bool
var count int

func bfs(start int) {
	queue := make([]int, 0)
	queue = append(queue, start)
	count++
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if visited[cur] == true {
			continue
		}
		visited[cur] = true
		for i := 1; i <= N; i++ {
			if grid[cur][i] == 1 && visited[i] == false {
				queue = append(queue, i)
			}
		}
	}
}

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bufout.Flush()
	fmt.Scanln(&N, &M)
	for i := 0; i < M; i++ {
		var start, end int
		fmt.Fscanln(bufin, &start, &end)
		grid[start][end] = 1
		grid[end][start] = 1
	}
	for j := 1; j <= N; j++ {
		if visited[j] == false {
			bfs(j)
		}
	}
	fmt.Fprintln(bufout, count)
}
