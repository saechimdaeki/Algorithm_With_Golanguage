package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, k int

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

var arr [101]int
var dynamic [10001]int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &k)
	for i := 1; i <= n; i++ {
		fmt.Fscan(bufin, &arr[i])
	}
	dynamic[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; (j + arr[i]) <= k; j++ {
			if dynamic[j] != 0 {
				if dynamic[j+arr[i]] == 0 {
					dynamic[j+arr[i]] = dynamic[j] + 1
				} else {
					dynamic[j+arr[i]] = min(dynamic[j+arr[i]], dynamic[j]+1)
				}
			}
		}
	}
	fmt.Fprint(bufout, dynamic[k]-1)
}
