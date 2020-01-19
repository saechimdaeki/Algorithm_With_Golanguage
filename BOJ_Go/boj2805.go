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
var vector [1000000]int64
var left, right, mid, solution int64

func max(x, y int64) int64 {
	if x > y {
		return x
	} else {
		return y
	}
}
func canI(x int64) int64 {
	var length int64
	length = 0
	for i := 0; i < n; i++ {
		if vector[i]-x > 0 {
			length += vector[i] - x
		}
	}
	if length >= m {
		return 123
	} else {
		return -1
	}
}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(bufin, &vector[i])

	}
	for i := 0; i < n; i++ {
		right = max(right, vector[i])
	}
	left = 1
	for {
		if left > right {
			break
		}

		mid = (left + right) / 2
		if canI(mid) >= 1 {
			solution = max(solution, mid)
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Fprint(bufout, solution)

}
