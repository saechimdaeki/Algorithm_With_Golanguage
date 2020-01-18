package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var n int
var m int64

func max(x, y int64) int64 {
	if x > y {
		return x
	} else {
		return y
	}
}
func min(x, y int64) int64 {
	if x > y {
		return y
	} else {
		return x
	}
}

var maximum int64

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	vector := make([]int64, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(bufin, &vector[i])
		maximum = max(maximum, vector[i])
	}
	fmt.Fscan(bufin, &m)
	var left, right, sum, result int64
	left = 0
	right = maximum
	for {
		if left > right {
			break
		}
		var mid int64
		sum = 0
		mid = (left + right) / 2
		for i := 0; i < n; i++ {
			sum += min(vector[i], mid)
			//fmt.Fprintln(bufout, sum)
		}
		if sum <= m {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	fmt.Fprint(bufout, result)

}
