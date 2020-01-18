package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var bx, by int
var dx, dy int
var jx, jy int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &bx, &by, &dx, &dy, &jx, &jy)
	cntb := max(abs(jx-bx), abs(jy-by))
	cntd := abs(jx-dx) + abs(jy-dy)
	if cntb < cntd {
		fmt.Fprint(bufout, "bessie")
	} else if cntb > cntd {
		fmt.Fprint(bufout, "daisy")

	} else {
		fmt.Fprint(bufout, "tie")

	}
}
