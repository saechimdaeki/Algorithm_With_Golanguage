package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}

}

var t, n int
var twocnt, fivecnt int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(bufin, &n)
		twocnt = 0
		fivecnt = 0
		for i := 2; i <= n; i *= 2 {
			twocnt += n / i
		}
		for i := 5; i <= n; i *= 5 {
			fivecnt += n / i
		}
		fmt.Fprint(bufout, Min(twocnt, fivecnt))
		fmt.Fprint(bufout, "\n")
	}
}
