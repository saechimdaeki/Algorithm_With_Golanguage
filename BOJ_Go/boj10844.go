package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int
var dp [101][10]int64
var ans int64

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	for i := 1; i <= 9; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 0; j <= 9; j++ {
			dp[i][j] = 0
			if j-1 >= 0 {
				dp[i][j] += dp[i-1][j-1]
			}
			if j+1 <= 9 {
				dp[i][j] += dp[i-1][j+1]
			}
			dp[i][j] %= 1000000000
		}
	}
	for i := 0; i <= 9; i++ {
		ans += dp[n][i]
	}
	fmt.Fprint(bufout, ans%1000000000)
}
