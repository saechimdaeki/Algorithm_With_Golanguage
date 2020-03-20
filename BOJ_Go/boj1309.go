package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int
var dp [100001][3]int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + dp[i-1][1] + dp[i-1][2]
		dp[i][1] = dp[i-1][0] + dp[i-1][2]
		dp[i][2] = dp[i-1][0] + dp[i-1][1]
		for j := 0; j < 3; j++ {
			dp[i][j] %= 9901
		}
	}
	fmt.Fprint(bufout, (dp[n][0]+dp[n][1]+dp[n][2])%9901)
}
