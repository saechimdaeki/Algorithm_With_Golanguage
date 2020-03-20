package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int
var dp [91][91]int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	dp[1][0] = 0
	dp[1][1] = 1
	for i := 2; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + dp[i-1][1]
		dp[i][1] = dp[i-1][0]
	}
	fmt.Fprint(bufout, dp[n][0]+dp[n][1])
}
