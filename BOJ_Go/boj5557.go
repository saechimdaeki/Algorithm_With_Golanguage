package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int
var arr [101]int
var dynamic [101][21]int64

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(bufin, &arr[i])

	}
	dynamic[1][arr[1]]++
	for i := 2; i < n; i++ {
		for j := 0; j <= 20; j++ {
			if dynamic[i-1][j] != 0 {
				if j+arr[i] <= 20 {
					dynamic[i][j+arr[i]] += dynamic[i-1][j]
				}
				if j-arr[i] >= 0 {
					dynamic[i][j-arr[i]] += dynamic[i-1][j]
				}
			}
		}
	}
	fmt.Fprint(bufout, dynamic[n-1][arr[n]])
}
