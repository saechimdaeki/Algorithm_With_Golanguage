package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var grid [101][101]int
var n, m, r int
var arr [101]int
var answer int

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &m, &r)
	for i := 1; i <= n; i++ {
		fmt.Fscan(bufin, &arr[i])
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i != j {
				grid[i][j] = 2e9 + 1
			}
		}
	}
	for i := 0; i < r; i++ {
		var tmp1, tmp2, tmp3 int
		fmt.Fscan(bufin, &tmp1, &tmp2, &tmp3)
		grid[tmp1][tmp2] = tmp3
		grid[tmp2][tmp1] = tmp3
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				if grid[j][k] > grid[j][i]+grid[i][k] {
					grid[j][k] = grid[j][i] + grid[i][k]
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		tmp := 0
		for j := 1; j <= n; j++ {
			if grid[i][j] <= m {
				tmp += arr[j]
			}
		}
		answer = max(tmp, answer)
	}
	fmt.Fprint(bufout, answer)
}
