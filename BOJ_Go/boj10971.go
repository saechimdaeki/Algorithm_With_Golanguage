package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

const maximum = 2e9 + 1

var solution int
var n int
var grid [11][11]int
var visited [11]bool

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func tsp(vertex, value, count int) {
	visited[1] = true
	if count == n-1 && grid[vertex][1] >= 1 {
		solution = min(solution, value+grid[vertex][1])
		return
	}
	for i := 1; i <= n; i++ {
		if grid[vertex][i] >= 1 && visited[i] == false {
			visited[i] = true
			tsp(i, value+grid[vertex][i], count+1)
			visited[i] = false
		}
	}
}
func main() {
	defer bufout.Flush()
	//fmt.Fprint(bufout, maximum)
	fmt.Fscan(bufin, &n)
	solution = maximum
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(bufin, &grid[i][j])
		}
	}
	tsp(1, 0, 0)
	fmt.Fprint(bufout, solution)
}
