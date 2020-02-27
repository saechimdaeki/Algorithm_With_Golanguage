package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var grid [101][101]int
var n int
var visited [101][101]int64

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(bufin, &grid[i][j])
		}
	}
	visited[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i+grid[i][j] < n {
				visited[i+grid[i][j]][j] += visited[i][j]
			}
			if j+grid[i][j] < n {
				visited[i][j+grid[i][j]] += visited[i][j]
			}
		}
	}
	fmt.Fprintln(bufout, visited[n-1][n-1])
}
