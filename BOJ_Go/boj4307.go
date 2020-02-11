package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var t int

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

var l, n int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(bufin, &l, &n)
		minimum := 0
		maximum := 0
		for j := 0; j < n; j++ {
			var tmp int
			fmt.Fscan(bufin, &tmp)
			A := min(tmp, l-tmp)
			B := max(tmp, l-tmp)
			minimum = max(minimum, A) //크흠;;
			maximum = max(maximum, B)
		}
		fmt.Fprintln(bufout, minimum, maximum)
	}
}
