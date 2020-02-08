package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var t, n int
var sum int

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(bufin, &n)
		sum = 0
		for i := 0; i < n; i++ {
			var a, b, c int
			fmt.Fscan(bufin, &a, &b, &c)
			if max(a, max(b, c)) > 0 {
				sum += max(a, max(b, c))
			}
		}
		fmt.Fprintln(bufout, sum)
	}
}
