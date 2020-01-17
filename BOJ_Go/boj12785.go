package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var w, h, x, y int
var dp [200][200]int64

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &w, &h)
	fmt.Fscan(bufin, &x, &y)
	for i := 0; i < h; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < w; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % 1000007
		}
	}
	fmt.Fprint(bufout, (dp[y-1][x-1]*dp[h-y][w-x])%1000007)
}
