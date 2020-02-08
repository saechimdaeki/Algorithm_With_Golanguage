package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var grid [501][501]int
var dynamic [501][501]int
var dir = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}
var n, m int

func search(x, y int) int {
	if x == m && y == n {
		return 1
	}
	if dynamic[x][y] == 12345 {
		dynamic[x][y] = 0
		for i := 0; i < 4; i++ {
			nextx := x + dir[i][0]
			nexty := y + dir[i][1]
			if nextx >= 1 && nextx <= m && nexty >= 1 && nexty <= n {
				if grid[x][y] > grid[nextx][nexty] {
					dynamic[x][y] += search(nextx, nexty)
				}

			}
		}
	}
	return dynamic[x][y]
}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &m, &n)

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			dynamic[i][j] = 12345
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(bufin, &grid[i][j])
		}
	}
	fmt.Fprint(bufout, search(1, 1))
}
