package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var t int
var array [1001]int
var time [1001]int
var dynamic [1001][1001]int
var n int

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func solve(n1 int) int {
	if array[n1] != 0 {
		return array[n1]
	} else {
		var answer int
		answer = 0
		for i := 1; i <= n; i++ {
			if dynamic[n1][i] == 1 {
				answer = max(answer, solve(i))
			}
		}
		array[n1] = answer + time[n1]
		return array[n1]
	}

}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &t)
	for t > 0 {
		var k, d int
		fmt.Fscan(bufin, &n, &k)
		for i := 1; i <= n; i++ {
			time[i] = 0
			array[i] = 0
			for j := 1; j <= n; j++ {
				dynamic[i][j] = 0
			}
		}

		for i := 1; i <= n; i++ {
			fmt.Fscan(bufin, &time[i])
		}

		for i := 0; i < k; i++ {
			var x, y int
			fmt.Fscan(bufin, &x, &y)
			dynamic[y][x] = 1
		}
		fmt.Fscan(bufin, &d)
		fmt.Fprint(bufout, solve(d))
		fmt.Fprint(bufout, "\n")
		t--
	}
}
