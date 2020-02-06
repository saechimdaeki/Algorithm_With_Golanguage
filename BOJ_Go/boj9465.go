package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var t, n int
var arr [2][100001]int
var dynamic [2][100001]int

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func solve(plz string) int {
	dynamic[0][0] = 0
	dynamic[1][0] = 0
	dynamic[0][1] = arr[0][1]
	dynamic[1][1] = arr[1][1]
	for i := 2; i <= n; i++ {
		dynamic[0][i] = max(dynamic[1][i-1], dynamic[1][i-2]) + arr[0][i]
		dynamic[1][i] = max(dynamic[0][i-1], dynamic[0][i-2]) + arr[1][i]
	}
	return max(dynamic[0][n], dynamic[1][n])

}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(bufin, &n)
		for j := 0; j < 2; j++ {
			for k := 1; k <= n; k++ {
				fmt.Fscan(bufin, &arr[j][k])
			}
		}

		fmt.Fprintln(bufout, solve("답을 알려줠 "))
	}
}
