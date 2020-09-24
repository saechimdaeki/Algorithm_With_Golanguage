package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var grid [501][501]int
var n, m int
var arr [501]int
var answer int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i != j {
				grid[i][j] = 1000000000
			}
		}
	}
	for i := 1; i <= m; i++ {
		var tmp1, tmp2 int
		fmt.Fscan(bufin, &tmp1, &tmp2)
		grid[tmp1][tmp2] = 1
	}
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if grid[i][j] > grid[i][k]+grid[k][j] {
					grid[i][j] = grid[i][k] + grid[k][j]
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				continue
			}
			if grid[i][j] < 1000000000 {
				arr[i]++
				arr[j]++
				if arr[i] == n-1 {
					answer++
				}
				if arr[j] == n-1 {
					answer++
				}
			}
		}
	}
	fmt.Fprint(bufout, answer)
}
