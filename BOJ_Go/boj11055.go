package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int
var arr [1000]int
var dp [1000]int

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
func dynamic() int {
	answer := 0
	for i := 0; i < n; i++ {
		dp[i] = arr[i]
		for j := 0; j <= i; j++ {
			if arr[i] > arr[j] && dp[i] < dp[j]+arr[i] {
				dp[i] = dp[j] + arr[i]
			}
		}
		answer = Max(answer, dp[i])

	}
	return answer
}

func main() {
	defer bufout.Flush()
	fmt.Fscanln(bufin, &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(bufin, "%d", &arr[i])
	}
	fmt.Fprintln(bufout, dynamic())
}
