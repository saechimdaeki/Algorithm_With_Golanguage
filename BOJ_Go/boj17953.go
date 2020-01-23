package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, m int

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

var maxi int
var arr [11][100001]int
var dynamic [11][100001]int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &m)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(bufin, &arr[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dynamic[j][i] = max(dynamic[j][i], dynamic[j][i-1]+(arr[j][i])/2)
			maxi = 0
			for k := 1; k <= m; k++ {
				if j == k {
					continue
				} else {
					maxi = max(maxi, dynamic[k][i-1])
				}
			}
			dynamic[j][i] = max(dynamic[j][i], maxi+arr[j][i])
		}
	}
	maxi = -1e9
	for i := 1; i <= m; i++ {
		maxi = max(maxi, dynamic[i][n])
	}
	fmt.Fprint(bufout, maxi)
}
