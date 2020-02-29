package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var grid [21][21]int
var n int
var visited [21][21]bool
var answer int

func floydwarshall() {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if j == k || k == i || j == i {
					continue
				} else if grid[j][k] > (grid[j][i] + grid[i][k]) {
					answer = 1
					return
				} else if grid[j][k] == (grid[j][i] + grid[i][k]) {
					visited[j][k] = false
				}
			}
		}
	}
}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(bufin, &grid[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			visited[i][j] = true
		}
	}
	floydwarshall()
	if answer == 1 {
		fmt.Fprint(bufout, "-1")
	} else {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if visited[i][j] == true {
					answer += grid[i][j]
				}
			}
		}
		fmt.Fprint(bufout, answer/2)
	}
}
