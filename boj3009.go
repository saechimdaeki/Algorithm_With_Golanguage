package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var x1, y1, x2, y2, x3, y3, x4, y4 int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &x1, &y1, &x2, &y2, &x3, &y3)
	if x1 == x2 {
		x4 = x3
	} else if x1 == x3 {
		x4 = x2
	} else {
		x4 = x1
	}
	if y1 == y2 {
		y4 = y3
	} else if y1 == y3 {
		y4 = y2
	} else {
		y4 = y1
	}
	fmt.Fprintf(bufout, "%d %d", x4, y4)
}
