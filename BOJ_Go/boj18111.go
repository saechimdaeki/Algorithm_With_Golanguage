package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, m, b int
var grid [501][501]int
var max int = -6e7
var min int = 6e7
var maximum int = 6e7
var answer int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &m, &b)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(bufin, &grid[i][j])
			if min > grid[i][j] {
				min = grid[i][j]
			}
			if max < grid[i][j] {
				max = grid[i][j]
			}
		}
	}
	for i := min; i <= max; i++ {
		t := 0
		cnt := b
		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				bts := grid[j][k] - i
				if bts > 0 {
					t = t + (2 * bts)
					cnt = cnt + bts
				} else if bts < 0 {
					t = t - bts
					cnt = cnt + bts
				}
			}
		}
		if cnt >= 0 {
			if t <= maximum {
				maximum = t
				answer = i
			}
		}
	}
	fmt.Fprintf(bufout, "%d %d", maximum, answer)
}
